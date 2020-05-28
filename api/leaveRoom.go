package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"net/http"
)


func LeaveRoom(ctx *fasthttp.RequestCtx) {
	roomName := ctx.UserValue("roomName").(string)

	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	}

	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		storageManager.DeleteUser(cookie, roomName, user)
		WriteResponse(ctx, http.StatusOK, model.ResponseMessage{model.OK})
		return
	} else {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoGetUserByCookie})
		return
	}
}