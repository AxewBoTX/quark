package lib

import (
	"os"

	"golang.org/x/net/websocket"
)

var DefaultConfig Config

var CurrentConfig Config

func Broadcaster() {
	for {
		select {
		case data := <-MSG_Channel:
			for _, client := range Clients {
				if message_broadcast_err := websocket.JSON.Send(client, data); message_broadcast_err != nil {
					ErrorWithColor(
						"ERROR",
						"0",
						COLOR_RED,
						"Failed To Broadcast Message",
						message_broadcast_err,
					)
				}
			}
		}
	}
}

func DisconnectClient(client_addr string) {
	delete(Clients, client_addr)
	InfoWithColor("LEAVE", "0", COLOR_YELLOW, "Client left the server", "Address", client_addr)
}

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
	if len(CurrentConfig.DatabaseFileName) == 0 || CurrentConfig.DatabaseFileName == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"DatabaseFileName value not found in config, using default value",
		)
		DB_FILE_PATH = SRC_FOLDER_PATH + "/" + DefaultConfig.DatabaseFileName
	} else {
		DB_FILE_PATH = SRC_FOLDER_PATH + "/" + CurrentConfig.DatabaseFileName
	}
	if len(CurrentConfig.UserTableName) == 0 || CurrentConfig.UserTableName == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"UserTableName value not found in config, using default value",
		)
		USER_TABLE_NAME = DefaultConfig.UserTableName
	} else {
		USER_TABLE_NAME = CurrentConfig.UserTableName
	}
	if len(CurrentConfig.MessageTableName) == 0 || CurrentConfig.MessageTableName == "" {
		WarnWithColor("WARN",
			"0",
			COLOR_YELLOW,
			"MessageTableName value not found in config, using default value",
		)
		MESSAGE_TABLE_NAME = DefaultConfig.MessageTableName
	} else {
		MESSAGE_TABLE_NAME = CurrentConfig.MessageTableName
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
