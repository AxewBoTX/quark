package lib

import (
	"database/sql"
	"fmt"

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

// create database tables
func HandleMigrations(DB *sql.DB) {
	if _, table_create_err := DB.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE,
		passwordHash TEXT,
		userAuthToken TEXT UNIQUE
	)`, USER_TABLE_NAME)); table_create_err != nil {
		log.Fatal("Failed To Create usr_base table", "Error", table_create_err)
	}
}
