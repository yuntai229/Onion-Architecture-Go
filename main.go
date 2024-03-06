package main

import (
	"net/http"
	cmd "onion-architecrure-go/cmd"
	"time"
)

func main() {
	db := cmd.InitDb()
	logger := cmd.InitLog()
	handlers, middlewares := cmd.InitApp(db, logger)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmd.InitRouter(handlers, middlewares),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
