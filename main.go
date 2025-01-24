package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"app/db"
	// "app/pkg"
	"app/routers"
)

func main() {
	println("Executando o servidor... \n")
	RunServerAPI()
}

func RunServerAPI() {
	// Carregar as variáveis de ambiente
	if err_load := godotenv.Load(); err_load != nil {
		fmt.Println("Erro no godotenv.Load")
		panic(err_load)
	}

	// Conexão com banco de dados
	var Conn, err = db.Conn_Postgres()
	defer Conn.Close()
	if err != nil {
		log.Fatal("Erro", err)
		return
	}

	// Framework Gin - http requests
	var server = gin.Default()

	var myConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	server.Use(cors.New(myConfig))

	// End Points
	routers.Init(server, Conn)

	server.Run(":8000")
}
