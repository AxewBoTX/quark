package lib

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/log"
	_ "modernc.org/sqlite"
)

// handle database file and folder creation and initialize SQL database
func CreateDatabase() *sql.DB {
	if CheckFileExists(DB_FILE_PATH) == false {
		CreateFile(DB_FILE_PATH)
	}
	DB, db_open_err := sql.Open("sqlite", DB_FILE_PATH)
	if db_open_err != nil {
		log.Fatal("Failed To Open Database", "Error", db_open_err)
	}
	return DB
}

var (
	TableCreateQuery = `CREATE TABLE IF NOT EXISTS %s(
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE,
		passwordHash TEXT,
		userAuthToken TEXT UNIQUE,
		created TEXT,
		updated TEXT
	);`
)

// create database tables
func HandleMigrations(DB *sql.DB) {
	if _, table_create_err := DB.Exec(fmt.Sprintf(TableCreateQuery, USER_TABLE_NAME)); table_create_err != nil {
		log.Fatal("Failed To Create usr_base table", "Error", table_create_err)
	}
}

// generate SQL UPDATE query based on User struct input
func GenerateSQLUpdateQuery(user User) string {
	current_time := time.Now().Format(time.RFC3339)
	var updateFields []string
	if user.Username != "" {
		updateFields = append(updateFields, fmt.Sprintf(`username='%s'`, user.Username))
	}
	if user.PasswordHash != "" {
		updateFields = append(updateFields, fmt.Sprintf(`passwordHash='%s'`, user.PasswordHash))
	}
	if user.UserAuthToken != "" {
		updateFields = append(updateFields, fmt.Sprintf(`userAuthToken='%s'`, user.UserAuthToken))
	}
	if len(updateFields) == 0 {
		return ""
	}
	updateFields = append(updateFields, fmt.Sprintf(`updated='%s'`, current_time))
	setClause := strings.Join(updateFields, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE ID='%s';", USER_TABLE_NAME, setClause, user.ID)
	return query
}
