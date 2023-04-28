package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/producer_rabbitmq/server"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic(err)
	}
}

func main() {
	server.Start()
}
