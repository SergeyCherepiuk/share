package http

import (
	"fmt"
	"net/http"

	"github.com/SergeyCherepiuk/share/server/pkg/http/ws"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/create", func(c echo.Context) error {
		websocket.Server{Handler: func(conn *websocket.Conn) {
			room := ws.NewRoom()
			go room.Start()

			fmt.Println(room.Id)

			room.Join(conn)
			defer room.Leave(conn)

			ListenForMessages(conn, room)
		}}.ServeHTTP(c.Response(), c.Request())
		return c.NoContent(http.StatusSwitchingProtocols)
	})

	e.GET("/join/:id", func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}

		room, err := ws.GetRoom(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		websocket.Server{Handler: func(conn *websocket.Conn) {
			room.Join(conn)
			defer room.Leave(conn)

			messages := make(chan string)
			defer close(messages)

			ListenForMessages(conn, room)
		}}.ServeHTTP(c.Response(), c.Request())
		return c.NoContent(http.StatusSwitchingProtocols)
	})

	return e
}

func ListenForMessages(conn *websocket.Conn, room *ws.Room) {
	var message string
	for {
		if err := websocket.Message.Receive(conn, &message); err != nil {
			return
		}
		room.Send(conn, message)
	}
}
