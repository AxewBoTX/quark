package lib

// constant declarations

const (
	PORT           string = ":8080"
	DB_FOLDER_PATH string = "database"
	DB_FILE_PATH   string = "database/base.db"
	COLOR_RED      string = "#f38ba8"
	COLOR_BLUE     string = "#89b4fa"
	COLOR_GREEN    string = "#a6e3a1"
	COLOR_YELLOW   string = "#f9e2af"
)

// type definitions
type (
	User struct {
		ID            string `json:"id"`
		Username      string `json:"username"`
		PasswordHash  string `json:"passwordHash"`
		UserAuthToken string `json:"userAuthToken"`
	}
)
