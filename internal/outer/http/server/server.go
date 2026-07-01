package server

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"app/internal/inner/usecase"
	"app/internal/outer/http/controller"
	"app/internal/outer/persistence/db"
	"app/internal/outer/persistence/repository"
)

func RunServer() {
	server := gin.Default()

	server.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: false,
			MaxAge:           12 * time.Hour,
		},
	))

	conn, err := db.ConnPostgreSQL()
	if err != nil {
		log.Println("[error ao conectar com o banco de dados]:", err)
		return
	}

	repo, err := repository.InitLayer(conn)
	if err != nil {
		log.Println("[error ao inicializar repositório]:", err)
		return
	}

	err = repo.CreateTable()
	if err != nil {
		log.Println("[error ao criar tabela]:", err)
		return
	}

	useCase := usecase.InitLayer(repo)
	controller := controller.InitLayer(useCase)

	Routers(server, controller)

	server.GET("/readiness", func(c *gin.Context) {
		if conn.Ping() != nil {
			c.JSON(503, gin.H{
				"postgres": false,
			})
			return
		}

		c.JSON(200, gin.H{
			"status":   "ready",
			"postgres": true,
		})
	})

	server.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8200"
	}

	err = server.Run(":" + port)
	if err != nil {
		log.Println("[error ao iniciar servidor]:", err)
		return
	}
}
