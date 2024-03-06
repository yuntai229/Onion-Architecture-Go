package cmd

import (
	"onion-architecrure-go/app"
	"onion-architecrure-go/infrastructure/rdb"
	"onion-architecrure-go/presentation/api/handler"
	"onion-architecrure-go/presentation/api/middleware"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitApp(db *gorm.DB, logger *zap.Logger) (handlers []any, middlewares []any) {
	userRepo := rdb.NewUserRepo(db, logger)
	threadRepo := rdb.NewThreadRepo(db, logger)

	userApp := app.NewUserApp(userRepo, logger)
	threadApp := app.NewThreadApp(threadRepo, logger)

	homeHandler := handler.NewHomeHandler(logger)
	userHandler := handler.NewUserHandler(userApp, logger)
	threadHandler := handler.NewThreadHandler(threadApp, logger)

	jwtMiddelware := middleware.NewJwtMiddleware(logger)
	logTraceMiddleware := middleware.NewLogTraceMiddleware(logger)

	handlers = []any{homeHandler, userHandler, threadHandler}
	middlewares = []any{jwtMiddelware, logTraceMiddleware}

	return handlers, middlewares
}
