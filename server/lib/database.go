package lib

import (
	"database/sql"
	"os"

	"github.com/charmbracelet/log"
	_ "modernc.org/sqlite"
)

// handle database file and folder creation and initialize SQL database
func PrepareDatabase() *sql.DB {
	if _, folder_check_err := os.Stat(SRC_FOLDER_PATH); os.IsNotExist(folder_check_err) {
		folder_create_err := os.Mkdir(SRC_FOLDER_PATH, 0755)
		if folder_create_err != nil {
			log.Fatal("Failed To Create Folder", "Error", folder_create_err)
		}
	}
	if _, file_check_err := os.Stat(DB_FILE_PATH); os.IsNotExist(file_check_err) {
		_, file_create_err := os.Create(DB_FILE_PATH)
		if file_create_err != nil {
			log.Fatal("Failed To Create DB File", "Error", file_create_err)
		}
	}
	DB, db_open_err := sql.Open("sqlite", DB_FILE_PATH)
	if db_open_err != nil {
		log.Fatal("Failed To Open Database", "Error", db_open_err)
	}
	return DB
}

// create database tables
func HandleMigrations(DB *sql.DB) {
	if _, table_create_err := DB.Exec(`CREATE TABLE IF NOT EXISTS usr_base(
		id TEXT PRIMARY KEY,
		username TEXT UNIQUE,
		passwordHash TEXT,
		userAuthToken TEXT UNIQUE
	)`); table_create_err != nil {
		log.Fatal("Failed To Create usr_base table", "Error", table_create_err)
	}
}
