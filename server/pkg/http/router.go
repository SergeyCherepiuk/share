package http

import (
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

			websocket.JSON.Send(conn, ws.Frame{
				Opcode:  ws.OPCODE_ROOM_INFO,
				Payload: []byte(room.Id.String()),
			})

			room.Join(conn, false)
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
			room.Join(conn, true)
			defer room.Leave(conn)

			ListenForMessages(conn, room)
		}}.ServeHTTP(c.Response(), c.Request())
		return c.NoContent(http.StatusSwitchingProtocols)
	})

	return e
}

func ListenForMessages(conn *websocket.Conn, room *ws.Room) {
	var frame ws.Frame
	for {
		if err := websocket.JSON.Receive(conn, &frame); err == nil {
			room.Send(conn, frame)
		}
	}
}
