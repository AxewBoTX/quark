package handlers

import (
	"errors"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"

	"quark/server/lib"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Quark Server")
}

func WebSocketHandler(c echo.Context) error {
	websocket.Handler(func(conn *websocket.Conn) {
		client_addr := c.Request().RemoteAddr
		lib.Clients[client_addr] = conn

		lib.InfoWithColor(
			"JOIN",
			"0",
			lib.COLOR_GREEN,
			"Client joined the server",
			"Address",
			client_addr,
		)

		defer func() {
			delete(lib.Clients, client_addr)
			conn.Close()
		}()

		for {
			var msg lib.Message
			if message_read_err := websocket.JSON.Receive(conn, &msg); message_read_err != nil {
				if errors.Is(message_read_err, io.EOF) {
					lib.DisconnectClient(client_addr)
					break
				} else {
					lib.DisconnectClient(client_addr)
					lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Failed To Read Client Message", "Error", message_read_err)
					break
				}
			}
			lib.MSG_Channel <- msg
			lib.InfoWithColor("MSG", "0", lib.COLOR_BLUE, client_addr, "Message", msg.Body)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
