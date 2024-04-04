package web

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"quark/client/lib"
)

func RenderTemplate(c echo.Context, status int, comp templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if template_render_err := comp.Render(c.Request().Context(), c.Response().Writer); template_render_err != nil {
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
