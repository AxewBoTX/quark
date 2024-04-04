package lib

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

const (
	PORT string = ":3000"
)

func RenderTemplate(c echo.Context, status int, comp templ.Component) error {
	c.Response().Writer.WriteHeader(status)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	if template_render_err := comp.Render(c.Request().Context(), c.Response().Writer); template_render_err != nil {
		return c.String(http.StatusInternalServerError, "Failed To Render Response Template")
	}
	return nil
}
