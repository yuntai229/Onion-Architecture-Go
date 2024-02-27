package cmd

import (
	"onion-architecrure-go/app"
	"onion-architecrure-go/infrastructure/rdb"
	"onion-architecrure-go/presentation/api/handler"

	"gorm.io/gorm"
)

func InitApp(db *gorm.DB) []any {
	userRepo := rdb.NewUserRepo(db)
	threadRepo := rdb.NewThreadRepo(db)

	userApp := app.NewUserApp(userRepo)
	threadApp := app.NewThreadApp(threadRepo)

	homeHandler := handler.NewHomeHandler()
	userHandler := handler.NewUserHandler(userApp)
	threadHandler := handler.NewThreadHandler(threadApp)

	handlers := []any{homeHandler, userHandler, threadHandler}

	return handlers
}
