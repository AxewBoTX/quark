package lib

import (
	"os"

	"golang.org/x/net/websocket"
)

var DefaultConfig Config

var CurrentConfig Config

// broadcast messages to all the connected clients
func Broadcaster() {
	for {
		select {
		case data := <-MSG_Channel:
			data.Log()
			for user_id, client := range Clients {
				var same bool
				if user_id == data.UserID {
					same = true
				} else {
					same = false
				}
				if message_broadcast_err := websocket.JSON.Send(client, map[string]interface{}{
					"id":       data.ID,
					"user_id":  "",
					"username": data.Username,
					"body":     data.Body,
					"type":     data.Type,
					"created":  data.Created,
					"same":     same,
				}); message_broadcast_err != nil {
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

// disconnect client by user struct
func DisconnectClient(user User) {
	delete(Clients, user.ID)
	MSG_Channel <- Message{UserID: user.ID, Username: user.Username, Type: "LEAVE"}
}

// prepare things before application startup
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
		CurrentConfig.DatabaseFileName,
		DefaultConfig.DatabaseFileName,
		"DatabaseFileName",
		func(val string) {
			DB_FILE_PATH = SRC_FOLDER_PATH + "/" + val
		},
	)
	checkAndSetConfig(
		CurrentConfig.UserTableName,
		DefaultConfig.UserTableName,
		"UserTableName",
		func(val string) {
			USER_TABLE_NAME = val
		},
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
