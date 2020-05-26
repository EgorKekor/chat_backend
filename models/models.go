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

type Room struct {
	Name              string
	Users			  map[string]*User
	HistoryRecord     []*HistoryRecord
}

func CreateRoom(name string) *Room {
	room := &Room {Name: name}
	return room
}

const OK = "Ok"
const CookieExist = "This cookie exist"
const NameDoublicate = "Doublicate name"
const NoCookie = "No cookie"
const BadFormat = "Bad message format"
const NoGetUserByCookie = "Impossible GetUserByCookie"
const ServerError = "Server error"
const UpgradeError = "UpgradeError"


type ResponseMessage struct {
	Message string `json:"message"`
}

const ResponseTypeRooms = "rooms"

type ResponseContent struct {
	Type string `json:"type"`
	Content map[string]interface{} `json:"content"`

}

type UserMessage struct {
	Text string `json:"text"`
}

