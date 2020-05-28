package models

type Void struct{}

type HistoryRecord struct {
	User *User
	UserName string
	Message *Message
}

type SerialisableHistoryRecord struct {
	User string			`json:"userName"`
	Message string	`json:"message"`
}

type User struct {
	Name string
	AllNames map[string][]*Room
	//CurrentRoom *Room
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
	//Users			  map[string]*User
	UsersCookie		  map[string]*User
	HistoryRecord     []*HistoryRecord
	NoWatched		  map[string][]*HistoryRecord	// Имя пользователя - непросмотренные сообщения
}

func CreateRoom(name string) *Room {
	room := &Room {Name: name}
	return room
}

const OK = "Ok"
const CookieExist = "This cookie exist"
const NameDoublicate = "Doublicate name"
const AlreadyExist = "Room already exist"
const RoomNotExist = "Room not exist"
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
	Content map[string]string `json:"content"`
}

type UserMessage struct {
	Text string `json:"text"`
}


type UpdateMessages struct {
	Type string											`json:"type"`
	Content	map[string][]SerialisableHistoryRecord		`json:"content"` 		// Имя комнаты - непрочитанные записи
}

type AllMessages struct {
	Type string								`json:"type"`
	Content	[]SerialisableHistoryRecord		`json:"content"` 		// Имя комнаты - непрочитанные записи
}













