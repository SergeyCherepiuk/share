package ws

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

var rooms = make(map[uuid.UUID]*Room)

type Room struct {
	Id uuid.UUID

	ctx         context.Context
	cancel      context.CancelFunc
	messages    chan Message
	connections map[*websocket.Conn]struct{}
}

func (r Room) Start() {
	for {
		select {
		case <-r.ctx.Done():
			return
		case message := <-r.messages:
			for conn := range r.connections {
				if conn != message.from {
					websocket.JSON.Send(conn, message)
				}
			}
		}
	}
}

func (r *Room) Join(conn *websocket.Conn) {
	r.connections[conn] = struct{}{}
}

func (r Room) Send(conn *websocket.Conn, msg string) {
	r.messages <- Message{
		Type: TYPE_MESSAGE,
		Msg:  msg,
		from: conn,
	}
}

func (r *Room) Leave(conn *websocket.Conn) {
	delete(r.connections, conn)
	if len(r.connections) == 0 {
		r.cancel()
		delete(rooms, r.Id)
	}
}

func NewRoom() *Room {
	ctx, cancel := context.WithCancel(context.Background())
	room := Room{
		Id:          uuid.New(),
		ctx:         ctx,
		cancel:      cancel,
		messages:    make(chan Message),
		connections: make(map[*websocket.Conn]struct{}),
	}
	rooms[room.Id] = &room
	return &room
}

func GetRoom(id uuid.UUID) (*Room, error) {
	room, ok := rooms[id]
	if !ok {
		return nil, fmt.Errorf("room not found")
	}

	return room, nil
}
