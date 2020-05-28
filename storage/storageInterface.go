package storage

import model "github.com/EgorKekor/chat_backend/models"

type Storage interface  {
	AddRoom(name string, room *model.Room)
	GetRoom(name string) (*model.Room, bool)
	AddUser(room *model.Room, userName, cookie string) string
	DeleteUser(cookie, roomName string, user *model.User)
	GetUserByCookie(cookie string) (*model.User, bool)
	AddMessage(userName, roomName, messageText string) bool
	GetHistory(room *model.Room) []model.SerialisableHistoryRecord
	GetRooms() (Rooms map[string]*model.Room)

	ReadHistory(room *model.Room, userName string)
}