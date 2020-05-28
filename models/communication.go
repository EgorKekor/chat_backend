package models

// Шлется в вебсокет
type UpdateMessages struct {
	Type string											`json:"type"`
	Content	map[string][]SerialisableHistoryRecord		`json:"content"` 		// Имя комнаты - непрочитанные записи
}

// Шлется при входе в комнату
type AllMessages struct {
	Type string											`json:"type"`
	Content	[]SerialisableHistoryRecord					`json:"content"` 		// Имя комнаты - непрочитанные записи
}

// Шлется всегда, когда не шлются данные, для удобства отладки на клиенте
type ResponseMessage struct {
	Message string 										`json:"message"`
}

// Пользователь присылает сообщение
type UserMessage struct {
	Text string 										`json:"text"`
}

// Список комнат на главной странице
const ResponseTypeRooms = "rooms"
type ResponseContent struct {
	Type string 										`json:"type"`
	Content map[string]string 							`json:"content"`
}

// Нерекурсивная запись истории
type SerialisableHistoryRecord struct {
	User string			`json:"userName"`
	Message string	`json:"message"`
}

const OK = "Ok"
const CookieExist = "This cookie exist"
const NameDoublicate = "Doublicate name"
const AlreadyLogined = "Already logined "
const AlreadyExist = "Room already exist"
const RoomNotExist = "Room not exist"
const NoCookie = "No cookie"
const BadFormat = "Bad message format"
const NoGetUserByCookie = "Impossible GetUserByCookie"
const ServerError = "Server error"
const UpgradeError = "UpgradeError"