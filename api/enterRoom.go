package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
	"net/http"
	"time"
)

func EnterRoom(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	roomName := ctx.UserValue("roomName").(string)
	userName := ctx.UserValue("userName").(string)

	var cookie string

	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		cookie = uuid.New().String()
		cookieObj := fasthttp.AcquireCookie()

		cookieObj.SetValue(cookie)
		cookieObj.SetExpire(time.Now().Add(512 * time.Hour))
		cookieObj.SetSecure(false)
		cookieObj.SetHTTPOnly(true)
		cookieObj.SetPath("/")

		ctx.Response.Header.SetCookie(cookieObj)
	} else {
		cookie = string(clientCookie)
	}

	var room *model.Room
	var ok bool
	if room, ok = storageManager.GetRoom(roomName); !ok {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.RoomNotExist})
	}

	if result := storageManager.AddUser(room, userName, cookie); result != model.OK {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{result})
		return
	} else {
		storageManager.ReadHistory(room, userName)

		response := model.AllMessages{}
		response.Type = "room_history"
		response.Content = storageManager.GetHistory(room)

		WriteResponse(ctx, http.StatusOK, response)
		return
	}
}