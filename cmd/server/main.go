package main

import (
	"log"

	"github.com/goodleby/pure-go-server/handlers"
	"github.com/goodleby/pure-go-server/server"
)

func main() {
	server, err := server.Create()
	if err != nil {
		log.Fatalf("error while creating a server: %s", err.Error())
	}

	server.AddRoute("/", handlers.RootHandler)
	server.AddRoute("/articles", handlers.ArticlesHandler)

	server.Start()
}
