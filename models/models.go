package models

type void struct{}

type HistoryRecord struct {
	User *User
	Message *Message
}

type User struct {
	Name string
	Room *Room
	Messages []*Message
}

type Message struct {
	Text string
	Owner* User
}

