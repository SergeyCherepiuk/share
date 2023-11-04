package websocket

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
			for ws := range r.connections {
				if ws != message.From {
					websocket.JSON.Send(ws, message)
				}
			}
		}
	}
}

func (r *Room) Join(ws *websocket.Conn) {
	r.connections[ws] = struct{}{}
}

func (r Room) Send(ws *websocket.Conn, msg string) {
	r.messages <- Message{
		Type: TYPE_MESSAGE,
		Msg:  msg,
		From: ws,
	}
}

func (r *Room) Leave(ws *websocket.Conn) {
	delete(r.connections, ws)
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
