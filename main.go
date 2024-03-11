package main

import (
	"net/http"
	cmd "onion-architecrure-go/cmd"
	"time"
)

func main() {
	config := cmd.InitAppEnv()
	db := cmd.InitDb(config)
	logger := cmd.InitLog()
	handlers, middlewares := cmd.InitApp(config, db, logger)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      cmd.InitRouter(handlers, middlewares),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
