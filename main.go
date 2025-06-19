package main

import (
	"app/internal/server"
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	// Carregar as variáveis de ambiente
	if err_load := godotenv.Load(); err_load != nil {
		fmt.Println("Erro no godotenv.Load")
		panic(err_load)
	}
}

func main() {
	server.RunServer()
}
