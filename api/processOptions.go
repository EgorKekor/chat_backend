package api

import (
	"github.com/valyala/fasthttp"
	"net/http"
)

func ProcessOptions(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.Response.Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response.Header.Set("Access-Control-Allow-Headers", "origin, content-type, content-length,accept")
	ctx.Response.Header.Set("Access-Control-Allow-Methods", "GET, PUT, POST, OPTIONS, DELETE")
	ctx.SetStatusCode(http.StatusOK)

	return
}