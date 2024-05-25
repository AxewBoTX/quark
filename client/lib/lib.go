package lib

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

var DefaultConfig Config

var CurrentConfig Config

// preparations before running the app process
func Prepare() {
	if CheckItemExists(SRC_FOLDER_PATH, true) == false {
		CreateItem(SRC_FOLDER_PATH, true)
	}
}

// handle configuration integration
func HandleConfig() {
	checkAndSetConfig := func(current, defaultVal, key string, assign func(string)) {
		if current == "" {
			WarnWithColor(
				"WARN",
				"0",
				COLOR_YELLOW,
				key+" value not found in config, using default value",
			)
			assign(defaultVal)
		} else {
			assign(current)
		}
	}

	checkAndSetConfig(
		CurrentConfig.Host,
		DefaultConfig.Host,
		"Host",
		func(val string) { HOST = val },
	)
	checkAndSetConfig(
		CurrentConfig.Port,
		DefaultConfig.Port,
		"Port",
		func(val string) { PORT = ":" + val },
	)
	checkAndSetConfig(
		CurrentConfig.ServerHost,
		DefaultConfig.ServerHost,
		"ServerHost",
		func(val string) { SERVER_HOST = val },
	)
	checkAndSetConfig(
		CurrentConfig.ServerPort,
		DefaultConfig.ServerPort,
		"ServerPort",
		func(val string) { SERVER_PORT = ":" + val },
	)
}

// functions for handling IO
func CheckItemExists(name string, isFolder bool) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			if isFolder {
				WarnWithColor(
					"WARN",
					"0",
					COLOR_YELLOW,
					"Folder does not exist",
					"FolderName",
					name,
					err,
				)
			} else {
				WarnWithColor("WARN", "0", COLOR_YELLOW, "File does not exist", "Filename", name, err)
			}
			return false
		}
		if isFolder {
			FatalWithColor(
				"FATAL",
				"0",
				COLOR_RED,
				"Failed to check if folder exists",
				"FolderName",
				name,
				err,
			)
		} else {
			FatalWithColor("FATAL", "0", COLOR_RED, "Failed to check if file exists", "Filename", name, err)
		}
		return false
	}
	return true
}

func CreateItem(name string, isFolder bool) {
	if isFolder {
		InfoWithColor("INFO", "0", COLOR_BLUE, "Creating New Folder", "FolderName", name)
		err := os.Mkdir(name, 0755)
		if err != nil {
			FatalWithColor(
				"FATAL",
				"0",
				COLOR_RED,
				"Failed To Create Folder",
				"FolderName",
				name,
				err,
			)
		}
	} else {
		InfoWithColor("INFO", "0", COLOR_BLUE, "Creating New File", "Filename", name)
		_, err := os.Create(name)
		if err != nil {
			FatalWithColor("FATAL", "0", COLOR_RED, "Failed To Create File", "Filename", name, err)
		}
	}
}

// hash input string
func HashString(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(bytes), err
}

// compare input string with existing hash
func CheckStringHash(input, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(input))
	return err == nil
}
