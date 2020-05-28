package api

import (
	st "github.com/EgorKekor/chat_backend/storage"
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

var storageManager st.Storage
var upgrader websocket.FastHTTPUpgrader

func init() {
	upgrader = websocket.FastHTTPUpgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		CheckOrigin: func(ctx *fasthttp.RequestCtx) bool {
			return true
		},
	}
	storageManager = st.CreateLocalStorage()
}






















