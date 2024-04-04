package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"quark/client/lib"
	"quark/client/web/routes"
)

func IndexHandler(c echo.Context) error {
	return lib.RenderTemplate(c, http.StatusOK, routes.Home_Page())
}

func RegisterHandler(c echo.Context) error {
	return lib.RenderTemplate(c, http.StatusOK, routes.Register_Page())
}

func LoginHandler(c echo.Context) error {
	return lib.RenderTemplate(c, http.StatusOK, routes.Login_Page())
}

func ChatHandler(c echo.Context) error {
	return lib.RenderTemplate(c, http.StatusOK, routes.Chat_Page())
}
