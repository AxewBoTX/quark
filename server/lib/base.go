package lib

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/pelletier/go-toml"
	"golang.org/x/net/websocket"
)

// constant declarations
const (
	SRC_FOLDER_PATH          string = "base_src"
	CURRENT_CONFIG_FILE_PATH string = SRC_FOLDER_PATH + "/" + "config.toml"
	DEFAULT_CONFIG_FILE_PATH string = "public/lib/config.toml"
	SESSION_COOKIE_NAME      string = "usr_session"
	COLOR_RED                string = "#f38ba8"
	COLOR_BLUE               string = "#89b4fa"
	COLOR_GREEN              string = "#a6e3a1"
	COLOR_YELLOW             string = "#f9e2af"
	COLOR_ROSEWATER          string = "#f5e0dc"
	COLOR_TEXT_DARK          string = "#1e1e2e"
	COLOR_TEXT_LIGHT         string = "#cdd6f4"
)

// variable declarations
var (
	HOST            string
	PORT            string
	DB_FILE_PATH    string
	USER_TABLE_NAME string
	Clients         = make(map[string]*websocket.Conn)
	MSG_Channel     = make(chan Message)
)

// type definitions
type (
	Config struct {
		Host             string `toml:"host"`
		Port             string `toml:"port"`
		DatabaseFileName string `toml:"database_file_name"`
		UserTableName    string `toml:"user_table_name"`
		MessageTableName string `toml:"message_table_name"`
	}
	User struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		PasswordHash  string `json:"passwordHash"`
		UserAuthToken string `json:"userAuthToken"`
		Created       int64  `json:"created"`
		Updated       int64  `json:"updated"`
	}
	Message struct {
		ID       string `json:"id"`
		UserID   string `json:"user_id"`
		Username string `json:"username"`
		Body     string `json:"body"`
		Type     string `json:"type"`
		Created  int64  `json:"created"`
	}
)

// load bytes into config
func (c *Config) LoadConfig(content []byte) {
	if config_parse_err := toml.Unmarshal(content, c); config_parse_err != nil {
		FatalWithColor("FATAL", "0", COLOR_RED, "Failed to parse config", "Error", config_parse_err)
	}
}

// log message to console
func (msg *Message) Log() {
	if msg.Type == "JOIN" {
		InfoWithColor(
			"JOIN",
			"0",
			COLOR_GREEN,
			msg.Username+" joined the server!",
			"User",
			map[string]interface{}{"UserID": msg.UserID, "Username": msg.Username},
		)
	} else if msg.Type == "MSG" {
		InfoWithColor("MSG", "0", COLOR_BLUE, map[string]interface{}{"UserID": msg.UserID, "Username": msg.Username}, "Body", msg.Body)
	} else if msg.Type == "LEAVE" {
		InfoWithColor(
			"LEAVE",
			"0",
			COLOR_YELLOW,
			msg.Username+" left the server!",
			"User",
			map[string]interface{}{"UserID": msg.UserID, "Username": msg.Username},
		)
	}
}

// custom log function
func InfoWithColor(heading, color_bg, color_fg string, message interface{}, opts ...interface{}) {
	styles := log.DefaultStyles()
	styles.Levels[log.InfoLevel] = lipgloss.NewStyle().
		SetString(heading).
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color(color_bg)).
		Foreground(lipgloss.Color(color_fg))
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
	})
	logger.SetStyles(styles)
	logger.Info(message, opts...)
}

func ErrorWithColor(heading, color_bg, color_fg string, message interface{}, opts ...interface{}) {
	styles := log.DefaultStyles()
	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString(heading).
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color(color_bg)).
		Foreground(lipgloss.Color(color_fg))
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
	})
	logger.SetStyles(styles)
	logger.Error(message, opts...)
}

func FatalWithColor(heading, color_bg, color_fg string, message interface{}, opts ...interface{}) {
	styles := log.DefaultStyles()
	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString(heading).
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color(color_bg)).
		Foreground(lipgloss.Color(color_fg))
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
	})
	logger.SetStyles(styles)
	logger.Fatal(message, opts...)
}

func WarnWithColor(heading, color_bg, color_fg string, message interface{}, opts ...interface{}) {
	styles := log.DefaultStyles()
	styles.Levels[log.WarnLevel] = lipgloss.NewStyle().
		SetString(heading).
		Padding(0, 1, 0, 1).
		Background(lipgloss.Color(color_bg)).
		Foreground(lipgloss.Color(color_fg))
	logger := log.NewWithOptions(os.Stdout, log.Options{
		ReportTimestamp: true,
	})
	logger.SetStyles(styles)
	logger.Warn(message, opts...)
}
