package api

import (
	model "github.com/EgorKekor/chat_backend/models"
	st "github.com/EgorKekor/chat_backend/storage"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"time"
)

func WsWorker(conn *websocket.Conn, storage st.Storage, room *model.Room) {
	defer conn.Close()
	for {
		records := storage.GetHistory(room)
		for record := range records {
			conn.Write();
		}
		time.Sleep(1 * time.Second)
	}
	return
}


func GetMessages(ctx *fasthttp.RequestCtx) {
	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	}

	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		room := user.Room
		err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
			go WsWorker(conn, storageManager, room)
		})
		if err != nil {
			log.Println("Upgrade:", err)
			WriteResponse(ctx, http.StatusInternalServerError, model.ResponseMessage{model.UpgradeError})
			return
		}
		return
	} else {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoGetUserByCookie})
		return
	}

}