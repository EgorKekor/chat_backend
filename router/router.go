package router

import (
	"github.com/EgorKekor/chat_backend/api"
	"github.com/buaazp/fasthttprouter"
)

func CreateRouter() *fasthttprouter.Router {
	rt := fasthttprouter.New()

	rt.POST("/api/room/create/:roomName", api.CreateRoom)
	rt.POST("/api/room/enter/:roomName/:userName", api.EnterRoom)
	rt.POST("/api/message/send/:roomName/:userName", api.AddMessage)
	rt.GET("/api/rooms", api.GetRooms)
	rt.GET("/api/ws", api.GetMessages)
	rt.GET("/api/room/read/:roomName", api.ReadMessages)

	rt.OPTIONS("/api/rooms", api.ProcessOptions)
	rt.OPTIONS("/api/room/create/:roomName", api.ProcessOptions)
	rt.OPTIONS("/api/room/enter/:roomName/:userName", api.ProcessOptions)
	rt.OPTIONS("/api/ws", api.ProcessOptions)
	rt.OPTIONS("/api/message/send/:roomName/:userName", api.ProcessOptions)
	rt.OPTIONS("/api/room/read/:roomName", api.ProcessOptions)

	return rt
}