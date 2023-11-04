package http

import (
	"fmt"
	"net/http"

	myws "github.com/SergeyCherepiuk/share/server/pkg/http/websocket"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func Router() *echo.Echo {
	e := echo.New()

	e.GET("/create", func(c echo.Context) error {
		websocket.Server{Handler: func(ws *websocket.Conn) {
			room := myws.NewRoom()
			go room.Start()

			fmt.Println(room.Id)

			room.Join(ws)
			defer room.Leave(ws)

			ListenForMessages(ws, room)
		}}.ServeHTTP(c.Response(), c.Request())
		return c.NoContent(http.StatusSwitchingProtocols)
	})

	e.GET("/join/:id", func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid id")
		}

		room, err := myws.GetRoom(id)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		websocket.Server{Handler: func(ws *websocket.Conn) {
			room.Join(ws)
			defer room.Leave(ws)

			messages := make(chan string)
			defer close(messages)

			ListenForMessages(ws, room)
		}}.ServeHTTP(c.Response(), c.Request())
		return c.NoContent(http.StatusSwitchingProtocols)
	})

	return e
}

func ListenForMessages(ws *websocket.Conn, room *myws.Room) {
	var message string
	for {
		if err := websocket.Message.Receive(ws, &message); err != nil {
			return
		}
		room.Send(ws, message)
	}
}
