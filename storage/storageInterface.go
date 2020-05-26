package storage

import model "github.com/EgorKekor/chat_backend/models"

type Storage interface  {
	AddRoom(name string, room *model.Room)
	GetRoom(name string) (*model.Room, bool)
	AddUser(room *model.Room, userName, cookie string) string
	DeleteUser(cookie string, user *model.User)
	GetUserByCookie(cookie string) (*model.User, bool)
	AddMessage(user *model.User, messageText string)
	GetHistory(room *model.Room) []*model.HistoryRecord
	GetRooms() (Rooms map[string]*model.Room)
}