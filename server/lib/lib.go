package lib

import (
	"os"

	"github.com/charmbracelet/log"
)

var DefaultConfig Config

var CurrentConfig Config

func Prepare() {
	if CheckFolderExists(SRC_FOLDER_PATH) == false {
		CreateFolder(SRC_FOLDER_PATH)
	}
}

func HandleConfig() {
	if len(CurrentConfig.Host) == 0 || CurrentConfig.Host == "" {
		log.Warn("Host value not found in config, using default value")
		HOST = DefaultConfig.Host
	} else {
		HOST = CurrentConfig.Host
	}
	if len(CurrentConfig.Port) == 0 || CurrentConfig.Port == "" {
		log.Warn("Port value not found in config, using default value")
		PORT = ":" + DefaultConfig.Port
	} else {
		PORT = ":" + CurrentConfig.Port
	}
	if len(CurrentConfig.DatabaseFileName) == 0 || CurrentConfig.DatabaseFileName == "" {
		log.Warn("DatabaseFileName value not found in config, using default value")
		DB_FILE_PATH = SRC_FOLDER_PATH + "/" + DefaultConfig.DatabaseFileName
	} else {
		DB_FILE_PATH = SRC_FOLDER_PATH + "/" + CurrentConfig.DatabaseFileName
	}
	if len(CurrentConfig.UserTableName) == 0 || CurrentConfig.UserTableName == "" {
		log.Warn("UserTableName value not found in config, using default value")
		USER_TABLE_NAME = DefaultConfig.UserTableName
	} else {
		USER_TABLE_NAME = CurrentConfig.UserTableName
	}
}

func CheckFileExists(filename string) bool {
	if _, file_check_err := os.Stat(filename); file_check_err != nil {
		if os.IsNotExist(file_check_err) {
			log.Warn("File does not exist", "Filename", filename, "Error", file_check_err)
			return false
		}
		log.Warn("Failed to check if file exists", "Filename", filename, "Error", file_check_err)
		return false
	}
	return true
}

func CheckFolderExists(foldername string) bool {
	if _, folder_check_err := os.Stat(foldername); folder_check_err != nil {
		if os.IsNotExist(folder_check_err) {
			log.Warn("Folder does not exist", "FolderName", foldername, "Error", folder_check_err)
			return false
		}
		log.Warn(
			"Failed to check if folder exists",
			"FolderName",
			foldername,
			"Error",
			folder_check_err,
		)
		return false
	}
	return true
}

func CreateFile(filename string) {
	log.Info("Creating New File", "Filename", filename)
	_, file_create_err := os.Create(filename)
	if file_create_err != nil {
		log.Warn("Failed To Create File", "Filename", filename, "Error", file_create_err)
	}
}

func CreateFolder(foldername string) {
	log.Info("Creating New Folder", "FolderName", foldername)
	folder_create_err := os.Mkdir(foldername, 0755)
	if folder_create_err != nil {
		log.Warn("Failed To Create Folder", "FolderName", foldername, "Error", folder_create_err)
	}
}
