package models

type Void struct{}

type HistoryRecord struct {
	User *User
	UserName string
	Message *Message
}

type User struct {
	Name string
	AllNames map[string][]*Room
	Rooms map[string]*Room
	Messages []*Message
}

type Message struct {
	Text string
	Owner* User
	Room* Room
}

type Room struct {
	Name              string
	Users			  map[string]*User				// Логины
	UsersCookie		  map[string]string				// Реальные юзеры каким логином зашли, логин навсегда для комнаты
	HistoryRecord     []*HistoryRecord
	NoWatched		  map[string][]*HistoryRecord	// Имя пользователя - непросмотренные сообщения
}

func CreateRoom(name string) *Room {
	room := &Room {Name: name}
	return room
}
