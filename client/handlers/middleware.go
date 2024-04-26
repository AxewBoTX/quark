package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"

	"quark/client/lib"
	"quark/client/web"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		client := resty.New()
		URL := strings.TrimSpace(c.Request().URL.String())
		// check if session cookie is present
		if !strings.HasPrefix(URL, "/auth") {
			session_cookie, session_cookie_get_err := c.Cookie(lib.SESSION_COOKIE_NAME)
			if session_cookie_get_err != nil { // Not Present
				// Check if user is on protected route
				if strings.HasPrefix(URL, "/chat") { // protected
					c.Redirect(http.StatusSeeOther, "/login")
				} else { // not protected
					web.RenderTemplTemplate(c, http.StatusOK)
				}
			} else { // Present
				res, user_fetch_err := client.R().Get(lib.SERVER_HOST + lib.SERVER_PORT + "/users/token/" + session_cookie.Value)
				if user_fetch_err != nil || res.StatusCode() == http.StatusInternalServerError { // Not A Valid Response
					// Check if user is on protected route
					if strings.HasPrefix(URL, "/chat") { // protected
						c.Redirect(http.StatusSeeOther, "/login")
					} else { // not protected
						web.RenderTemplTemplate(c, http.StatusOK)
					}
				} else { // Valid Response
					// check if user is on protected route
					if strings.HasPrefix(URL, "/chat") {
						var user lib.User
						if resp_decode_err := json.Unmarshal(res.Body(), &user); resp_decode_err != nil {
							lib.ErrorWithColor(
								"ERROR",
								"0",
								lib.COLOR_RED,
								"Failed To Decode Server Response",
								"Error",
								resp_decode_err,
							)
						}
						c.Set("session-data", user)
						web.RenderTemplTemplate(c, http.StatusOK)
					} else { // not protected
						c.Redirect(http.StatusSeeOther, "/chat")
					}
				}
			}
		}
		return next(c)
	}
}
