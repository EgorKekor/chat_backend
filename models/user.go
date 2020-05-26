package models

type User struct {
	Cookie string
	Name string
	Room *Room
	Messages []*Message
}

func (us *User) addMessage(message *Message) {
	us.Messages = append(us.Messages, message)
}