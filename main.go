package main

import (
	"net/http"
	"onion-architecrure-go/app"
	cmd "onion-architecrure-go/cmd"
	"onion-architecrure-go/infrastructure/rdb"
	"time"
)

func main() {

	db := cmd.InitDb()

	userRepo := rdb.NewUserRepo(db)
	threadRepo := rdb.NewThreadRepo(db)

	userApp := app.NewUserApp(userRepo)
	threadApp := app.NewThreadApp(threadRepo)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmd.InitRouter(userApp, threadApp),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
