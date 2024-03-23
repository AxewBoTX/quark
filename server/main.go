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
	//basic application setup with default config
	lib.Prepare()
	default_config_bytes, default_config_load_err := PublicDir.ReadFile(
		lib.DEFAULT_CONFIG_FILE_PATH,
	)
	if default_config_load_err != nil {
		log.Fatal("Failed to load default config file", "Error", default_config_load_err)
	}
	lib.DefaultConfig.LoadConfig(default_config_bytes)

	// current config setup
	if lib.CheckFileExists(lib.CURRENT_CONFIG_FILE_PATH) == true {
		current_config_bytes, config_file_read_err := os.ReadFile(lib.CURRENT_CONFIG_FILE_PATH)
		if config_file_read_err != nil {
			log.Fatal("Failed to read current config file", "Error", config_file_read_err)
		}
		lib.CurrentConfig.LoadConfig(current_config_bytes)
	} else {
		if _, config_file_create_err := os.Create(lib.CURRENT_CONFIG_FILE_PATH); config_file_create_err != nil {
			log.Fatal("Failed to create config file", "Error", config_file_create_err)
		}
		if config_write_err := os.WriteFile(lib.CURRENT_CONFIG_FILE_PATH, default_config_bytes, 0644); config_write_err != nil {
			log.Fatal("Failed to write to config file", "Error", config_write_err)
		}
		lib.CurrentConfig.LoadConfig(default_config_bytes)
	}
	lib.HandleConfig()

	// database preparations
	DB := lib.CreateDatabase()
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
