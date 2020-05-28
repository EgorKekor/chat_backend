package api

import (
	"encoding/json"
	"fmt"
	model "github.com/EgorKekor/chat_backend/models"
	st "github.com/EgorKekor/chat_backend/storage"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
	"time"
)


func WsWorker(conn *websocket.Conn, storage st.Storage, user *model.User) {
	defer conn.Close()
	rooms := user.Rooms
	ticker := time.NewTicker(2 * time.Second)
	for {
		if err := conn.SetWriteDeadline(time.Now().Add(10 * time.Second)); err != nil {
			break
		}
		conn.SetCloseHandler(func(code int, text string) error {
			fmt.Printf("Closed-code:%d\ntext:%s")
			return nil
		})

		responseHistory := model.UpdateMessages{}
		responseHistory.Type = "update_messages"

		responseHistory.Content = make(map[string][]model.SerialisableHistoryRecord)
		for roomName, room := range rooms {
			responseHistory[roomName] = model.SerialisableHistoryRecord{room.NoWatched[user.Name], room.NoWatched[user.Messages]}
		}

		//for _, room := range rooms {
		//	for _, record := range room.HistoryRecord {
		//		responseHistory.Content = append(responseHistory.Content, model.SerialisableHistoryRecord{User: record.UserName, Message: record.Message.Text})
		//	}
		//}


		resp, err := json.Marshal(responseHistory)
		if err != nil {
			ticker.Stop()
			break
		}

		if err := conn.WriteMessage(websocket.TextMessage, resp); err != nil {
			fmt.Print("Error sending event: ", err)
			return
		}

		<-ticker.C
	}
	return
}


func GetMessages(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")

	var cookie string
	if clientCookie := ctx.Request.Header.Peek("Cookie"); clientCookie == nil {
		WriteResponse(ctx, http.StatusForbidden, model.ResponseMessage{model.NoCookie})
		return
	} else {
		cookie = string(clientCookie)
	}


	if user, ok := storageManager.GetUserByCookie(cookie); ok {
		err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
			WsWorker(conn, storageManager, user)
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