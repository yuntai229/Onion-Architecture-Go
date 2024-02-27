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
	handlers := cmd.InitApp(db)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmd.InitRouter(handlers),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
