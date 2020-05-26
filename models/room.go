package models

import st "github.com/EgorKekor/chat_backend/storage"


type Room struct {
	Name              string
	Users			  map[string]*User
	HistoryRecord     []*HistoryRecord
	storageManagerPtr *st.LocalStorage
}

func CreateRoom(name string, db *st.LocalStorage) *Room {
	room := &Room {Name: name}
	room.storageManagerPtr = db
	return room
}





