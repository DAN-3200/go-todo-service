package main

import (
	"log"

	"github.com/joho/godotenv"

	"app/internal/outer/http/server"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("[.env error]: ", err)
		return
	}
}

func main() {
	server.RunServer()
}
