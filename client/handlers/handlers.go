package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// index(/) route handler
func IndexHandler(c echo.Context) error {
	return nil
}

// (/register) route handler
func RegisterHandler(c echo.Context) error {
	return nil
}

// (/login) route handler
func LoginHandler(c echo.Context) error {
	return nil
}

// (/chat) route handler
func ChatHandler(c echo.Context) error {
	fmt.Println(c.Get("session-data"))
	return nil
}
