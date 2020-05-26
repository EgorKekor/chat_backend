package storage

import m "github.com/EgorKekor/chat_backend/models"

type Storage interface  {
	addUser(cookie, name string) bool
	getUser(name string) *m.User
	deleteUser(name string) bool
}