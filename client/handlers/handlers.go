package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"quark/client/web"
	"quark/client/web/routes"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK, routes.Home_Page())
}

// (/register) route handler
func RegisterHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK, routes.Register_Page())
}

// (/login) route handler
func LoginHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK, routes.Login_Page())
}

// (/chat) route handler
func ChatHandler(c echo.Context) error {
	return web.RenderTemplate(c, http.StatusOK, routes.Chat_Page())
}
