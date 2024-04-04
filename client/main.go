package main

import (
	"embed"
	"errors"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"quark/client/handlers"
	"quark/client/lib"
)

//go:embed all:public/lib
var LibDir embed.FS

func main() {
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true
	log.Info("Client started successfully", "PORT", ":3000")

	server.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: http.FS(LibDir),
		Root:       "/",
	}))

	server.GET("/", handlers.IndexHandler)
	server.GET("/register", handlers.RegisterHandler)
	server.GET("/login", handlers.LoginHandler)
	server.GET("/chat", handlers.ChatHandler)
	server.POST("/auth/login", handlers.AuthLoginHandler)
	server.POST("/auth/register", handlers.AuthRegisterHandler)

	if server_close_err := server.Start(lib.PORT); server_close_err != nil &&
		!errors.Is(server_close_err, http.ErrServerClosed) {
		log.Fatal("Server Closed", "Error", server_close_err)
	}
}
