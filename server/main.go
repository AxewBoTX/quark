package main

import (
	"embed"
	"errors"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"quark/server/handlers"
	"quark/server/lib"
)

//go:embed all:public/lib
var LibDir embed.FS

func main() {
	//basic application setup with default config
	lib.Prepare()
	default_config_bytes, default_config_load_err := LibDir.ReadFile(
		lib.DEFAULT_CONFIG_FILE_PATH,
	)
	if default_config_load_err != nil {
		lib.FatalWithColor(
			"FATAL",
			"0",
			lib.COLOR_RED,
			"Failed to load default config file",
			"Error",
			default_config_load_err,
		)
	}
	lib.DefaultConfig.LoadConfig(default_config_bytes)

	// current config setup
	if lib.CheckFileExists(lib.CURRENT_CONFIG_FILE_PATH) == true {
		current_config_bytes, config_file_read_err := os.ReadFile(lib.CURRENT_CONFIG_FILE_PATH)
		if config_file_read_err != nil {
			lib.FatalWithColor(
				"FATAL",
				"0",
				lib.COLOR_RED,
				"Failed to read config file",
				"Error",
				config_file_read_err,
			)
		}
		lib.CurrentConfig.LoadConfig(current_config_bytes)
	} else {
		lib.CreateFile(lib.CURRENT_CONFIG_FILE_PATH)
		if config_write_err := os.WriteFile(lib.CURRENT_CONFIG_FILE_PATH, default_config_bytes, 0644); config_write_err != nil {
			lib.FatalWithColor("FATAL", "0", lib.COLOR_RED, "Failed to write to config file", "Error", config_write_err)
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

	lib.InfoWithColor(
		"INFO",
		"0",
		lib.COLOR_BLUE,
		"Quark server started successfully",
		"URL",
		"http://localhost"+lib.PORT,
	)

	// route handlers
	server.GET("/", handlers.IndexHandler)
	handlers.Users(server.Group("/users"), DB)
	handlers.Messages(server.Group("/messages"), DB)

	// this function runs after the main function has ended
	defer func() {
		DB.Close()
		server.Close()
		lib.InfoWithColor(
			"INFO",
			"0",
			lib.COLOR_YELLOW,
			"Quark server closed successfully",
		)
	}()

	// starting echo web server
	if server_close_err := server.Start(lib.PORT); server_close_err != nil &&
		!errors.Is(server_close_err, http.ErrServerClosed) {
		lib.FatalWithColor("FATAL", "0", lib.COLOR_RED, "Server Closed", "Error", server_close_err)
	}
}
