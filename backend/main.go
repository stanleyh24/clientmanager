package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/stanleyh24/clientmanager/api"
	"github.com/stanleyh24/clientmanager/storage"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	store, err := storage.NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(*store)

	server.Run()
}
