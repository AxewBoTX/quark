package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Quark Server")
}
