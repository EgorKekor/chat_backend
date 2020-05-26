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

	if ok = storageManager.AddUser(room, userName, cookie); !ok {
		return false
	}
	return true
}


func addMessage(cookie, message string) bool {
	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		storageManager.AddMessage(user, message)
		return true
	} else {
		return false
	}
}


func LeaveRoom(cookie string) bool {
	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		storageManager.DeleteUser(cookie, user)
		return true
	} else {
		return false
	}
}


func GetMessages(cookie string) ([]*model.HistoryRecord, bool) {
	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		return storageManager.GetHistory(user.Room), true
	} else {
		return nil, false
	}
}





















