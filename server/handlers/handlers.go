package handlers

import (
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
	websocket.Handler(func(c *websocket.Conn) {
		lib.InfoWithColor("INFO", "0", lib.COLOR_BLUE, "New Connection!")
		client_session_cookie, client_session_cookie_get_err := c.Request().
			Cookie(lib.SESSION_COOKIE_NAME)
		if client_session_cookie_get_err != nil {
			lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Failed To Fetch Client Session Cookie")
		} else {
			lib.InfoWithColor("INFO", "0", lib.COLOR_BLUE, "Session Cookie Found!", "Value", client_session_cookie.Value)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
