package api

import (
	"fmt"
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"net/http"
)

func CreateRoom(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	roomName := ctx.UserValue("roomName").(string)

	if room, ok := storageManager.GetRoom(roomName); !ok {
		room = model.CreateRoom(roomName)
		storageManager.AddRoom(roomName, room)
		WriteResponse(ctx, http.StatusOK, model.ResponseMessage{model.OK})
		fmt.Printf("Created room:%s\n", roomName)
		return
	}

	WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.AlreadyExist})
	fmt.Printf("Create failed\n")
	return
}