package main

import (
	"net/http"
	"onion-architecrure-go/app"
	api "onion-architecrure-go/presentation/api"
	"time"
)

func main() {

	userApp := app.NewUser()
	userApp := app.NewUserApp()

	server := &http.Server{
		Addr:         ":8080",
		Handler:      api.InitRouter(userApp),
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	server.ListenAndServe()
}
