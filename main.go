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
	userApp := app.NewUserApp(userRepo)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmd.InitRouter(userApp),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
