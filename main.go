package main

import (
	"github.com/EgorKekor/chat_backend/router"
	"github.com/valyala/fasthttp"
)

func main() {
	rt := router.CreateRouter()
	fasthttp.ListenAndServe(":8081", rt.Handler)
}