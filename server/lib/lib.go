package lib

import (
	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml"
)

// constant declarations
const (
	PORT            string = ":8080"
	SRC_FOLDER_PATH string = "base_src"
	DB_FILE_PATH    string = SRC_FOLDER_PATH + "/" + "base.db"
	COLOR_RED       string = "#f38ba8"
	COLOR_BLUE      string = "#89b4fa"
	COLOR_GREEN     string = "#a6e3a1"
	COLOR_YELLOW    string = "#f9e2af"
)

// type definitions
type (
	Config struct {
		Host             string `toml:"host"`
		Port             string `toml:"port"`
		DatabaseFileName string `toml:"database_file_name"`
		UserTableName    string `toml:"user_table_name"`
	}
	User struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		PasswordHash  string `json:"passwordHash"`
		UserAuthToken string `json:"userAuthToken"`
	}
)

// load bytes into config
func (c *Config) LoadConfig(content []byte) {
	if config_parse_err := toml.Unmarshal(content, c); config_parse_err != nil {
		log.Fatal("Failed to parse config", "Error", config_parse_err)
	}
}
