package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"net/http"
	"strconv"
)


func GetRooms(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	rooms := storageManager.GetRooms()

	responseContent := model.ResponseContent{}
	responseContent.Type = model.ResponseTypeRooms
	responseContent.Content = map[string]string{}

	for _, room := range rooms {
		responseContent.Content[room.Name] = strconv.Itoa(len(room.Users))
	}
	WriteResponse(ctx, http.StatusOK, responseContent)

	return
}