package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"net/http"
)

func EnterRoom(ctx *fasthttp.RequestCtx) {
	roomName := ctx.UserValue("roomName").(string)
	userName := ctx.UserValue("userName").(string)

	var cookie string

	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		cookie = uuid.New().String()
		cookieObj := fasthttp.AcquireCookie()
		cookieObj.SetSecure(true)
		cookieObj.SetValue(cookie)
		ctx.Response.Header.SetCookie(cookieObj)
	}

	var room *model.Room
	var ok bool
	if room, ok = storageManager.GetRoom(roomName); !ok {
		room = model.CreateRoom(roomName)
		storageManager.AddRoom(roomName, room)
	}

	code := http.StatusOK
	var result string
	if result = storageManager.AddUser(room, userName, cookie); result != model.OK {
		code = http.StatusForbidden
	}

	WriteResponse(ctx, code, model.ResponseMessage{result})
	return
}