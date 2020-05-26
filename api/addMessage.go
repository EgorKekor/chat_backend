package api

import (
	"encoding/json"
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func AddMessage(ctx *fasthttp.RequestCtx) {
	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	}

	var userMessage model.UserMessage
	body := ctx.PostBody()
	if err := json.Unmarshal(body, &userMessage); err != nil {
		log.Println(err.Error())
		WriteResponse(ctx, http.StatusBadRequest, model.ResponseMessage{model.BadFormat})
	}

	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		storageManager.AddMessage(user, userMessage.Text)
		WriteResponse(ctx, http.StatusOK, model.ResponseMessage{model.OK})
		return
	} else {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoGetUserByCookie})
	}
}