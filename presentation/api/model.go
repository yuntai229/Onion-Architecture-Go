package api

import (
	"onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"
)

type Handlers struct {
	HomeHandler   *handler.HomeHandler
	UserHandler   *handler.UserHandler
	ThreadHandler *handler.ThreadHandler
}

type Middlewares struct {
	JwtMiddelware      *middleware.JwtAuthMiddleware
	LogTraceMiddleware *middleware.LogTraceMiddleware
}
