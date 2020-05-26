package api

import (
	"encoding/json"
	model "github.com/EgorKekor/chat_backend/models"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func GetRooms(ctx *fasthttp.RequestCtx) {
	rooms := storageManager.GetRooms()

	responseContent := model.ResponseContent{}
	responseContent.Type = model.ResponseTypeRooms
	responseContent.Content = map[string]interface{}{}

	responseContent.Content["a"] = 1
	for _, room := range rooms {
		responseContent.Content[room.Name] = len(room.Users)
	}

	if jsonObject, err := json.Marshal(&responseContent); err != nil {
		log.Println(err.Error())
		WriteResponse(ctx, http.StatusInternalServerError, model.ResponseMessage{model.ServerError})
	} else {
		stringify := string(jsonObject)
		WriteResponse(ctx, http.StatusOK, stringify)
	}

	return
}