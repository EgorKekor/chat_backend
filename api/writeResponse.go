package api

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

func WriteResponse(ctx *fasthttp.RequestCtx, statusCode int, body interface{}) {
	ctx.SetContentType("application/json")

	resp, err := json.Marshal(body)
	if err != nil {
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	if _, err := ctx.Write(resp); err != nil {
		log.Println(err.Error())
		ctx.SetStatusCode(http.StatusInternalServerError)
		return
	}
	ctx.SetStatusCode(statusCode)
}