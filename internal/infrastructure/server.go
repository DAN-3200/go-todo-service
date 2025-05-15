package infrastructure

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"app/internal/controller"
	"app/internal/db"
	"app/internal/repository"
	"app/internal/usecase"
)

func RunServerAPI() {
	// Carregar as variáveis de ambiente
	if err_load := godotenv.Load(); err_load != nil {
		fmt.Println("Erro no godotenv.Load")
		panic(err_load)
	}

	// Conexão com banco de dados
	var Conn = db.Conn_Sqlite()
	defer Conn.Close()
	// CreateTable(Conn)

	// `Arquitetura Limpa` em Camadas
	var useControllers = controller.Init(
		usecase.Init(
			repository.Init(Conn),
		),
	)

	// Framework Gin - http requests
	var server = gin.Default()

	var myCORSConfig = cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
	server.Use(cors.New(myCORSConfig))

	// End Points
	Routers(server, useControllers)

	server.Run(":8000")
}

func CreateTable(database *sql.DB) {
	var _, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS ToDo (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Title TEXT,
			Content TEXT,
			Status Boolean
		);`,
	)

	if err != nil {
		fmt.Println("Erro", err)
	}
}
