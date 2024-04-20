package web

import (
	"net/http"
	"strings"

	"github.com/a-h/templ"
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
	URL := strings.TrimSpace(c.Request().URL.String())

	_, session_cookie_get_err := c.Cookie(lib.SESSION_COOKIE_NAME)
	if session_cookie_get_err != nil { // if session cookie is not present
		if strings.HasPrefix(URL, "/chat") {
			// redirecting the the user to /login route if the user is current at any protected route
			c.Redirect(http.StatusSeeOther, "/login")
		} else {
			//rendering web page templ component according to current URL
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
		}
	} else { // if session cookie is present
		if URL == "/" || URL == "/login" || URL == "/register" {
			// redirecting the the user to /chat route if the user is current at any public unprotected route
			c.Redirect(http.StatusSeeOther, "/chat")
		} else {
			//rendering web page templ component according to current URL
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
		}
	}
	return nil
}
