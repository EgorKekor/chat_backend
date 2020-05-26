package api

import st "github.com/EgorKekor/chat_backend/storage"
import model "github.com/EgorKekor/chat_backend/models"

var storageManager *st.LocalStorage

func init() {
	storageManager = st.CreateStorage()
}

func EnterRoom(roomName, userName, cookie string) bool {		// Подключиться или создать
	var room *model.Room
	var ok bool
	if room, ok = storageManager.GetRoom(roomName); !ok {
		room = model.CreateRoom(roomName, storageManager)
		storageManager.AddRoom(roomName, room)
	}

	if ok = room.AddUser(userName, cookie); !ok {
		return false
		// return bad message
	}
	return true
	// return good message
}


func addMessage(userName, cookie, message string) bool {
	if room, ok := storageManager.GetRoomByUser(userName, cookie); ok {
		room.AddMessage(userName, cookie, message);
	}
}




















