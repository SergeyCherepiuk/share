package ws

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/SergeyCherepiuk/share/client/pkg/ws"
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type (
	RoomNotFound error
	RoomIsEmpty  error
)

var (
	muRooms sync.RWMutex
	rooms   = make(map[uuid.UUID]*Room)
)

type Room struct {
	Id uuid.UUID

	ctx    context.Context
	cancel context.CancelFunc

	frames chan Frame

	muConnection sync.RWMutex
	connections  map[*websocket.Conn]time.Time

	muWaitingConnections sync.RWMutex
	waitingConnections   map[*websocket.Conn]struct{}
}

func NewRoom() *Room {
	ctx, cancel := context.WithCancel(context.Background())
	room := Room{
		Id:                 uuid.New(),
		ctx:                ctx,
		cancel:             cancel,
		frames:             make(chan Frame),
		connections:        make(map[*websocket.Conn]time.Time),
		waitingConnections: make(map[*websocket.Conn]struct{}),
	}

	muRooms.Lock()
	rooms[room.Id] = &room
	muRooms.Unlock()

	go room.start()

	return &room
}

func GetRoom(id uuid.UUID) (*Room, error) {
	muRooms.RLock()
	room, ok := rooms[id]
	muRooms.RUnlock()
	if !ok {
		return nil, RoomNotFound(fmt.Errorf("room with id '%s' isn't found", id.String()))
	}

	return room, nil
}

func (r *Room) start() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case frame := <-r.frames:
			r.handleFrame(frame)
		}
	}
}

func (r *Room) Join(conn *websocket.Conn, fetchContent bool) {
	r.muConnection.Lock()
	r.connections[conn] = time.Now()
	r.muConnection.Unlock()

	if fetchContent {
		r.muWaitingConnections.Lock()
		r.waitingConnections[conn] = struct{}{}
		r.muWaitingConnections.Unlock()

		if oldestConn, err := r.getOldestConnection(); err == nil {
			websocket.JSON.Send(oldestConn, ws.Frame{
				Opcode: ws.OPCODE_CONTENT_REQUEST,
			})
		}
	}
}

func (r *Room) Send(conn *websocket.Conn, frame Frame) {
	frame.from = conn
	r.frames <- frame
}

func (r *Room) Leave(conn *websocket.Conn) {
	r.muConnection.Lock()
	defer r.muConnection.Unlock()

	delete(r.connections, conn)
	if len(r.connections) == 0 {
		r.cancel()
		muRooms.Lock()
		delete(rooms, r.Id)
		muRooms.Unlock()
	}
}

func (r *Room) getOldestConnection() (*websocket.Conn, error) {
	var (
		oldestConn       *websocket.Conn
		earliestJoinTime = time.Now()
	)

	r.muConnection.RLock()
	for conn, joinTime := range r.connections {
		if joinTime.Before(earliestJoinTime) {
			oldestConn = conn
			earliestJoinTime = joinTime
		}
	}
	r.muConnection.RUnlock()

	if oldestConn == nil {
		return nil, RoomIsEmpty(fmt.Errorf("room is empty"))
	}

	return oldestConn, nil
}

func (r *Room) handleFrame(frame Frame) {
	switch frame.Opcode {
	case OPCODE_OPERATIONS:
		r.muConnection.RLock()
		for conn := range r.connections {
			if conn != frame.from {
				websocket.JSON.Send(conn, frame)
			}
		}
		r.muConnection.RUnlock()

	case OPCODE_CONTENT_RESPONSE:
		r.muWaitingConnections.Lock()
		for conn := range r.waitingConnections {
			websocket.JSON.Send(conn, frame)
		}
		r.waitingConnections = make(map[*websocket.Conn]struct{})
		r.muWaitingConnections.Unlock()
	}
}
