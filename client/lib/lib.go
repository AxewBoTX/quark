package lib

import (
	"os"
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
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"Host value not found in config, using default value",
		)
		HOST = DefaultConfig.Host
	} else {
		HOST = CurrentConfig.Host
	}
	if len(CurrentConfig.Port) == 0 || CurrentConfig.Port == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"Port value not found in config, using default value",
		)
		PORT = ":" + DefaultConfig.Port
	} else {
		PORT = ":" + CurrentConfig.Port
	}
	if len(CurrentConfig.ServerHost) == 0 || CurrentConfig.ServerHost == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"ServerHost value not found in config, using default value",
		)
		SERVER_HOST = DefaultConfig.ServerHost
	} else {
		SERVER_HOST = CurrentConfig.ServerHost
	}
	if len(CurrentConfig.ServerPort) == 0 || CurrentConfig.ServerPort == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"ServerPort value not found in config, using default value",
		)
		SERVER_PORT = ":" + DefaultConfig.ServerPort
	} else {
		SERVER_PORT = ":" + CurrentConfig.ServerPort
	}
}

func CheckFileExists(filename string) bool {
	if _, file_check_err := os.Stat(filename); file_check_err != nil {
		if os.IsNotExist(file_check_err) {
			WarnWithColor("WARN",
				"0",
				COLOR_YELLOW,
				"File does not exist",
				"Filname",
				filename,
				"Error",
				file_check_err,
			)
			return false
		}
		FatalWithColor("FATAL",
			"0",
			COLOR_RED,
			"Failed to check if file exists",
			"Filname",
			filename,
			"Error",
			file_check_err,
		)
		return false
	}
	return true
}

func CheckFolderExists(foldername string) bool {
	if _, folder_check_err := os.Stat(foldername); folder_check_err != nil {
		if os.IsNotExist(folder_check_err) {
			WarnWithColor("WARN",
				"0",
				COLOR_YELLOW,
				"Folder does not exist",
				"FolderName",
				foldername,
				"Error",
				folder_check_err,
			)
			return false
		}
		FatalWithColor("FATAL",
			"0",
			COLOR_RED,
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
	InfoWithColor(
		"INFO",
		"0",
		COLOR_BLUE,
		"Creating New File",
		"Filename",
		filename,
	)
	_, file_create_err := os.Create(filename)
	if file_create_err != nil {
		FatalWithColor(
			"FATAL",
			"0",
			COLOR_RED,
			"Failed To Create File",
			"Filename",
			filename,
			"Error",
			file_create_err,
		)
	}
}

func CreateFolder(foldername string) {
	InfoWithColor(
		"INFO",
		"0",
		COLOR_BLUE,
		"Creating New Folder",
		"FolderName",
		foldername,
	)
	folder_create_err := os.Mkdir(foldername, 0755)
	if folder_create_err != nil {
		FatalWithColor(
			"FATAL",
			"0",
			COLOR_RED,
			"Failed To Create Folder",
			"FolderName",
			foldername,
			"Error",
			folder_create_err,
		)
	}
}
