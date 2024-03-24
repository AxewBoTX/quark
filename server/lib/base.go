package lib

import (
	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml"
)

// constant declarations
const (
	SRC_FOLDER_PATH          string = "base_src"
	CURRENT_CONFIG_FILE_PATH string = SRC_FOLDER_PATH + "/" + "config.toml"
	DEFAULT_CONFIG_FILE_PATH string = "public/lib/config.toml"
	COLOR_RED                string = "#f38ba8"
	COLOR_BLUE               string = "#89b4fa"
	COLOR_GREEN              string = "#a6e3a1"
	COLOR_YELLOW             string = "#f9e2af"
)

// variable declarations
var (
	HOST            string
	PORT            string
	DB_FILE_PATH    string
	USER_TABLE_NAME string
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
		Created       string `json:"created"`
		Updated       string `json:"updated"`
	}
)

// load bytes into config
func (c *Config) LoadConfig(content []byte) {
	if config_parse_err := toml.Unmarshal(content, c); config_parse_err != nil {
		log.Fatal("Failed to parse config", "Error", config_parse_err)
	}
}
