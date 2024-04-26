package web

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/go-resty/resty/v2"
	"github.com/labstack/echo/v4"

	"quark/client/lib"
	"quark/client/web/routes"
)

// web pages
var Pages = map[string]templ.Component{
	"/":         routes.Home_Page(),
	"/register": routes.Register_Page(),
	"/login":    routes.Login_Page(),
	"/chat":     routes.Chat_Page(),
}

func RenderTemplate(c echo.Context, status int) error {
	client := resty.New()
	URL := strings.TrimSpace(c.Request().URL.String())
	// check if session cookie is present
	session_cookie, session_cookie_get_err := c.Cookie(lib.SESSION_COOKIE_NAME)
	if session_cookie_get_err != nil { // Not Present
		// Check if user is on protected route
		if strings.HasPrefix(URL, "/chat") { // protected
			c.Redirect(http.StatusSeeOther, "/login")
		} else { // not protected
			RenderTemplTemplate(c, http.StatusOK)
		}
	} else { // Present
		res, user_fetch_err := client.R().Get(lib.SERVER_HOST + lib.SERVER_PORT + "/users/token/" + session_cookie.Value)
		if user_fetch_err != nil || res.StatusCode() == http.StatusInternalServerError { // Not A Valid Response
			// Check if user is on protected route
			if strings.HasPrefix(URL, "/chat") { // protected
				c.Redirect(http.StatusSeeOther, "/login")
			} else { // not protected
				RenderTemplTemplate(c, http.StatusOK)
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
				c.Set("session-val", user)
				RenderTemplTemplate(c, http.StatusOK)
			} else { // not protected
				c.Redirect(http.StatusSeeOther, "/chat")
			}
		}
	}
	return nil
}

func RenderTemplTemplate(c echo.Context, status int) error {
	URL := strings.TrimSpace(c.Request().URL.String())
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if template_render_err := Pages[URL].Render(c.Request().Context(), c.Response().Writer); template_render_err != nil {
		lib.ErrorWithColor(
			"ERROR",
			"0",
			lib.COLOR_RED,
			"Failed To Render Response Template",
			"Error",
			template_render_err,
		)
		return c.String(http.StatusInternalServerError, "Failed To Render Response Template")
	}
	return nil
}
