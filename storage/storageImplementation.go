package storage

import (
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
	ls.Rooms[name].UsersCookie = make(map[string]string)
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

	if login, exist := room.UsersCookie[cookie]; exist { // Если заходил под каким либо логином
		if login != userName {
			return login //Не пустить под новым логином
		} else {
			return model.OK // Пустить под старым логином
		}
	}

	if _, dublicate := room.Users[userName]; dublicate {
		return model.NameDoublicate
	}

	if !currentExist {
		userPtr = &model.User{
			Name: userName,
			Rooms: make(map[string]*model.Room),
			AllNames: make(map[string][]*model.Room),
			Messages: make([]*model.Message, 0)}
		ls.Cookie[cookie] = userPtr
	}

	userPtr.Rooms[room.Name] = room
	room.Users[userName] = userPtr
	room.UsersCookie[cookie] = userName
	userPtr.AllNames[userName] = append(userPtr.AllNames[userName], room)


	return model.OK

}

func (ls *LocalStorage) ReadHistory(room *model.Room, userName string) {
	if _, exist := room.NoWatched[userName]; exist {
		room.NoWatched[userName] = room.NoWatched[userName][:0]
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

func (ls *LocalStorage) GetUserLoginInRoomByCookie(roomName, cookie string) (string, bool) {
	if r, existR := ls.Rooms[roomName]; existR {
		if login, existL := r.UsersCookie[cookie]; existL {
			return login, true
		}
		return "", false
	}
	return "", false
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












