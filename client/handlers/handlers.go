package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"quark/client/web"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK)
}

// (/register) route handler
func RegisterHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK)
}

// (/login) route handler
func LoginHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK)
}

// (/chat) route handler
func ChatHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK)
}
