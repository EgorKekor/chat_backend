package storage

import model "github.com/EgorKekor/chat_backend/models"

type roomEntities struct {
	Messages map[string]*model.Message
	Users map[string]*model.User
}

type LocalStorage struct {
	Rooms map[string]*model.Room
	Entities map[string]roomEntities
}

func CreateStorage() *LocalStorage {
	return &LocalStorage {
		Rooms: make(map[string]*model.Room),
		Entities: make(map[string]roomEntities),
	}
}

func (ls *LocalStorage) AddRoom(name string, room *model.Room) {
	ls.Rooms[name] = room
	ls.Entities[name] = roomEntities{Users: make(map[string]*model.User), Messages: make(map[string]*model.Message)}
}

func (ls *LocalStorage) GetRoom(name string) (*model.Room, bool) {
	r, b := ls.Rooms[name]
	return r, b
}

func (ls *LocalStorage) AddUser(roomName string, user *model.User) bool {
	if ent, ok := ls.Entities[roomName]; ok {				// Комната существует?
		if user, exist := ent.Users[user.Name]; exist {		// Пользователь существует?
			return false
		} else {
			ent.Users[user.Name] = user						// Создать пользователя
			return true
		}
	} else {
		return false
	}
}

func (ls *LocalStorage) GetRoomByUser(userName, cookie string) (*model.Room, bool) {
	for roomName, ent := range ls.Entities {		// Ищем во всех комнатах
		if user, ok := ent.Users[userName]; ok {	// Если нашли
			if user.Cookie == cookie {				// Нужно проверить по куке, так как имена могут совпадать
				return ls.Rooms[roomName], true		// Вернуть адрес комнаты
			}
		}
	}
	return nil, false
}

func (ls *LocalStorage) GetUserByRN(roomName, userName string) (*model.User, bool) {
	if ent, ok := ls.Entities[roomName]; ok {
		user, exist := ent.Users[userName]
		return user, exist
	}
	return nil, false
}


func (ls *LocalStorage) AddMessage(roomName, userName string) (*model.User, bool) {
	if ent, ok := ls.Entities[roomName]; ok {
		if messages, exist := ent.Messages[userName]; exist {

		}
	}
	return nil, false
}











