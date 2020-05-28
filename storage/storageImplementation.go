package storage

import (
	"fmt"
	model "github.com/EgorKekor/chat_backend/models"
)


type LocalStorage struct {
	Rooms map[string]*model.Room
	Cookie map[string]*model.User
}

func CreateLocalStorage() *LocalStorage {
	return &LocalStorage {
		Rooms: make(map[string]*model.Room),
		Cookie: make(map[string]*model.User),
	}
}


func (ls *LocalStorage) AddRoom(name string, room *model.Room) {
	ls.Rooms[name] = room
	ls.Rooms[name].Users = make(map[string]*model.User)
	ls.Rooms[name].NoWatched = make(map[string][]*model.HistoryRecord)
}

func (ls *LocalStorage) GetRoom(name string) (*model.Room, bool) {
	r, b := ls.Rooms[name]
	return r, b
}

func (ls *LocalStorage) GetRooms() (Rooms map[string]*model.Room) {
	return ls.Rooms
}


func (ls *LocalStorage) AddUser(room *model.Room, userName, cookie string) string {
	userPtr, currentExist := ls.Cookie[cookie]

	if currentExist {
		if rooms, logined := userPtr.AllNames[userName]; logined {			// Уже заходил под таким именем
			for _, visitedRoom := range rooms {								// Проверить все комнаты куда он под этим именем заходил
				if visitedRoom == room {									// Если повторно заходит
					return model.OK											// Тогда ОК
				}
			}
		}
	}

	if dublicateUser, dublicateExist := room.Users[userName]; dublicateExist {
		if currentExist && (userPtr != dublicateUser) {
			return model.NameDoublicate
		} else if currentExist {
			fmt.Printf("Repeat enter cookie:%s  userName:%s\n")
			return model.OK
		} else {
			return model.NameDoublicate
		}
	}

	if currentExist {
		userPtr.Rooms[room.Name] = room
		room.Users[userName] = userPtr
		userPtr.AllNames = make(map[string]*model.Room)
		userPtr.AllNames[userName] = room
	} else {
		newUserPtr := &model.User{Name: userName, Rooms: make(map[string]*model.Room, 0), Messages: make([]*model.Message, 0)}
		ls.Cookie[cookie] = newUserPtr
		newUserPtr.Rooms[room.Name] = room
		room.Users[userName] = newUserPtr
		newUserPtr.AllNames = make(map[string]*model.Room)
		newUserPtr.AllNames[userName] = room
	}

	return model.OK

}

func (ls *LocalStorage) ReadHistory(room *model.Room, userName string) {
	if userHistoryList, exist := room.NoWatched[userName]; exist {
		userHistoryList = userHistoryList[:0]
	}
}


func (ls *LocalStorage) DeleteUser(cookie, roomName string, user *model.User) {
	delete(user.Rooms[roomName].Users, user.Name)		// Удалить пользователя из комнаты roomName
	delete(user.Rooms, roomName)						// Удалить комнату roomName из списка комнат пользователя
}


func (ls *LocalStorage) GetUserByCookie(cookie string) (*model.User, bool) {
	if c, exist := ls.Cookie[cookie]; exist {
		return c, true
	}
	return nil, false
}



func (ls *LocalStorage) AddMessage(userName, roomName, messageText string) bool {
	var room *model.Room
	var ok bool
	if room, ok = ls.GetRoom(roomName); !ok {
		return false
	}

	var user *model.User
	if user, ok = room.Users[userName]; !ok {
		return false
	}


	newMessagePtr := &model.Message{Owner: user, Text: messageText, Room: room}
	user.Messages = append(user.Messages, newMessagePtr)

	record := &model.HistoryRecord{user, userName, newMessagePtr}
	for otherUserName, _ := range room.Users {
		if otherUserName != user.Name {
			room.NoWatched[otherUserName] = append(room.NoWatched[otherUserName], record)
		}
	}
	room.HistoryRecord = append(room.HistoryRecord, record)
	return true

}

func (ls *LocalStorage) GetHistory(room *model.Room) []model.SerialisableHistoryRecord {
	records := make([]model.SerialisableHistoryRecord, 0)

	for _, record := range room.HistoryRecord {
		records = append(records, model.SerialisableHistoryRecord{User: record.UserName, Message: record.Message.Text})
	}
	return records
}












