package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func initHttpServer()  {
	router := fasthttprouter.New()
	router.GET("/", Index)

	log.Fatal().Err(fasthttp.ListenAndServe(":8080", router.Handler))
}

func Index(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome!")
}
