package storage

import model "github.com/EgorKekor/chat_backend/models"


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
}

func (ls *LocalStorage) GetRoom(name string) (*model.Room, bool) {
	r, b := ls.Rooms[name]
	return r, b
}

func (ls *LocalStorage) GetRooms() (Rooms map[string]*model.Room) {
	return ls.Rooms
}

func (ls *LocalStorage) AddUser(room *model.Room, userName, cookie string) string {
	if _, exist := ls.Cookie[cookie]; exist {
		return model.CookieExist
	}

	if _, exist := room.Users[userName]; exist {
		return model.NameDoublicate
	} else {
		newUserPtr := &model.User{Name: userName, Room: room, Messages: make([]*model.Message, 0)}
		ls.Cookie[cookie] = newUserPtr
		room.Users[userName] = newUserPtr
		return model.OK
	}
}


func (ls *LocalStorage) DeleteUser(cookie string, user *model.User) {
	delete(user.Room.Users, user.Name)
	delete(ls.Cookie, cookie)
}


func (ls *LocalStorage) GetUserByCookie(cookie string) (*model.User, bool) {
	if c, exist := ls.Cookie[cookie]; exist {
		return c, true
	}
	return nil, false
}


func (ls *LocalStorage) AddMessage(user *model.User, messageText string) {
	newMessagePtr := &model.Message{Owner: user, Text: messageText}
	user.Messages = append(user.Messages, newMessagePtr)				// Добавили юзеру ссылку на сообщение
	user.Room.HistoryRecord = append(
		user.Room.HistoryRecord,
		&model.HistoryRecord{user, newMessagePtr})			// Добавили навседа запись в историю комнаты
}

func (ls *LocalStorage) GetHistory(room *model.Room) []*model.HistoryRecord {
	return room.HistoryRecord
}












