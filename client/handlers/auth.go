package handlers

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func AuthRegisterHandler(c echo.Context) error {
	log.Info(
		map[string]interface{}{
			"username":        c.FormValue("username"),
			"password":        c.FormValue("password"),
			"passwordConfirm": c.FormValue("passwordConfirm"),
		},
	)
	return c.JSON(http.StatusOK, "PASS")
}

func AuthLoginHandler(c echo.Context) error {
	log.Info(
		map[string]interface{}{
			"username": c.FormValue("username"),
			"password": c.FormValue("password"),
		},
	)
	return c.JSON(http.StatusOK, "PASS")
}
