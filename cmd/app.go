package cmd

import (
	"onion-architecrure-go/app"
	"onion-architecrure-go/domain/model"
	"onion-architecrure-go/infrastructure/rdb"
	"onion-architecrure-go/presentation/api"
	"onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitApp(config *model.Config, db *gorm.DB, logger *zap.Logger) (handlers api.Handlers, middlewares api.Middlewares) {
	userRepo := rdb.NewUserRepo(db, logger)
	threadRepo := rdb.NewThreadRepo(db, logger)

	userApp := app.NewUserApp(config, userRepo, logger)
	threadApp := app.NewThreadApp(threadRepo, logger)

	handlers = api.Handlers{
		HomeHandler:   handler.NewHomeHandler(logger),
		UserHandler:   handler.NewUserHandler(userApp, logger),
		ThreadHandler: handler.NewThreadHandler(threadApp, logger),
	}

	middlewares = api.Middlewares{
		JwtMiddelware:      middleware.NewJwtMiddleware(config, logger),
		LogTraceMiddleware: middleware.NewLogTraceMiddleware(logger),
	}

	return handlers, middlewares
}
