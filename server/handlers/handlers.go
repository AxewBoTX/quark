package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/go-resty/resty/v2"
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
		rest_client := resty.New()
		// get session cookie from the client connection
		session_cookie, session_cookie_get_err := c.Request().Cookie(lib.SESSION_COOKIE_NAME)
		if session_cookie_get_err != nil {
			lib.FatalWithColor(
				"FATAL",
				"0",
				lib.COLOR_RED,
				"Failed To Fetch Session Cookie",
				"Error",
				session_cookie_get_err,
			)
		}
		// fetch user data from the database according to the session cookie
		res, user_fetch_err := rest_client.R().
			Get("http://" + lib.HOST + lib.PORT + "/users/token/" + session_cookie.Value)
		if user_fetch_err != nil {
			lib.FatalWithColor(
				"FATAL",
				"0",
				lib.COLOR_RED,
				"Failed To Fetch User Related To Session Cookie",
				"Error",
				user_fetch_err,
			)
		}
		// decode server JSON response
		var user lib.User
		if resp_decode_err := json.Unmarshal(res.Body(), &user); resp_decode_err != nil {
			lib.FatalWithColor(
				"ERROR",
				"0",
				lib.COLOR_RED,
				"Failed To Decode Server Response",
				"Error",
				resp_decode_err,
			)
		}

		// append client to clients array
		lib.Clients[user.ID] = conn

		// broadcast client join message
		lib.MSG_Channel <- lib.Message{UserID: user.ID, Username: user.Username, Type: "JOIN"}

		defer func() {
			delete(lib.Clients, user.ID)
			conn.Close()
		}()

		// Message read loop
		for {
			var msg lib.Message
			if message_read_err := websocket.JSON.Receive(conn, &msg); message_read_err != nil {
				if errors.Is(message_read_err, io.EOF) {
					lib.DisconnectClient(user)
					break
				} else {
					lib.DisconnectClient(user)
					lib.ErrorWithColor("ERROR", "0", lib.COLOR_RED, "Failed To Read Client Message", "Error", message_read_err)
					break
				}
			}
			msg.UserID = user.ID
			msg.Username = user.Username
			msg.Type = "MSG"
			lib.MSG_Channel <- msg // broadcast client message
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
