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

// render the templ HTMX template
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
