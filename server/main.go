package main

import (
	"embed"
	"errors"
	"net/http"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	"quark/server/handlers"
	"quark/server/lib"
)

//go:embed all:public/lib
var PublicDir embed.FS

func main() {
	// database preparations
	DB := lib.PrepareDatabase()
	lib.HandleMigrations(DB)

	// echo web server
	server := echo.New()
	server.HideBanner = true
	server.HidePort = true

	log.Info("Quark server started successfully", "URL", "http://localhost"+lib.PORT)

	// route handlers
	server.GET("/", handlers.IndexHandler)
	handlers.Users(server.Group("/users"), DB)

	// this function runs after the main function has ended
	defer func() {
		DB.Close()
		server.Close()
		styles := log.DefaultStyles()
		styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
			Padding(0, 1, 0, 1).
			Background(lipgloss.Color(lib.COLOR_YELLOW)).
			Foreground(lipgloss.Color("0"))
		logger := log.New(os.Stdout)
		logger.SetStyles(styles)
		logger.Info("Quark server closed successfully")
	}()

	// starting echo web server
	if server_close_err := server.Start(lib.PORT); server_close_err != nil &&
		!errors.Is(server_close_err, http.ErrServerClosed) {
		log.Fatal("Server Closed", "Error", server_close_err)
	}
}
