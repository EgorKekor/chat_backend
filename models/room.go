package models

import st "github.com/EgorKekor/chat_backend/storage"

// HistoryRecord - история ссобщений, вынесена в модель из соображений разгрузки БД, т.к.
// При реальной реализации БД, в нее стоит сохранять вместе с сообщением дату его отправки.
// С помощью этой информации можно будет с легкостью восстановить историю из базы при потере
// Данных из оперативной памяти (1 запрос в бд) и далее вести локальную запись, продолжая
// Дублировать в БД.
type Room struct {
	Name              string
	Users			  []User
	HistoryRecord     HistoryRecord
	storageManagerPtr *st.LocalStorage
}

func CreateRoom(name string, db *st.LocalStorage) *Room {
	room := &Room {Name: name}
	room.storageManagerPtr = db
	return room
}

func (r *Room) AddUser(userName, cookie string) bool {
	newUserPtr := &User{Cookie: cookie, Name: userName, Room: r, Messages: make([]*Message, 0)}
	return r.storageManagerPtr.AddUser(r.Name, newUserPtr)
}

func (r *Room) AddMessage(userName, cookie, messageText string) bool {
	if user, ok := r.storageManagerPtr.GetUserByRN(r.Name, userName); ok {
		newMessagePtr := &Message{Owner: user, Text: messageText}
		user.Messages = append(user.Messages, newMessagePtr)
	}

	return false
}

