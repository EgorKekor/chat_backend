package router

import (
	"github.com/EgorKekor/chat_backend/api"
	"github.com/buaazp/fasthttprouter"
)

func CreateRouter() *fasthttprouter.Router {
	rt := fasthttprouter.New()

	rt.POST("/api/room/enter/:roomName/:userName", api.EnterRoom)
	rt.GET("/api/message/send", api.AddMessage)
	rt.DELETE("/api/room/leave", api.LeaveRoom)
	rt.GET("/api/rooms", api.GetRooms)
	rt.GET("/api/ws", api.GetMessages)

	return rt
}