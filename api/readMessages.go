package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"net/http"
)

func ReadMessages(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	roomName := ctx.UserValue("roomName").(string)

	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	} else {
		cookie = string(clientCookie)
	}


	if room, rOk := storageManager.GetRoom(roomName); rOk {
		if userLogin, uOk := storageManager.GetUserLoginInRoomByCookie(roomName, cookie); uOk {
			storageManager.ReadHistory(room, userLogin)

			WriteResponse(ctx, http.StatusOK, model.OK)
			return
		}
	}

	WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.RoomNotExist})
}