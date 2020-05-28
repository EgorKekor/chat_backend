package api

import (
	"encoding/json"
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func AddMessage(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	roomName := ctx.UserValue("roomName").(string)
	userName := ctx.UserValue("userName").(string)

	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	} else {
		cookie = string(clientCookie)
	}

	var userMessage model.UserMessage
	body := ctx.PostBody()
	if err := json.Unmarshal(body, &userMessage); err != nil {
		log.Println(err.Error())
		WriteResponse(ctx, http.StatusBadRequest, model.ResponseMessage{model.BadFormat})
	}


	storageManager.AddMessage(userName, roomName, userMessage.Text)

	if room, rOk := storageManager.GetRoom(roomName); rOk {
		if user, uOk := storageManager.GetUserByCookie(cookie); uOk {
			storageManager.ReadHistory(room, user.Name)

			response := model.AllMessages{}
			response.Type = "room_history"
			response.Content = storageManager.GetHistory(room)
			WriteResponse(ctx, http.StatusOK, response)
			return
		}
	}

	WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.RoomNotExist})

}