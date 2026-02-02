package server

import (
	"app/internal/outer/http/controller"
	"app/internal/outer/persistence/db"
	"app/internal/outer/persistence/repository"
	"app/internal/inner/usecase"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

	conn := db.ConnPostgreSQL()
	repo, err := repository.InitLayer(conn)
	if err != nil {
		log.Fatal("Error de resposta", err)
	}
	repo.CreateTable()
	useCase := usecase.InitLayer(repo)
	controller := controller.InitLayer(useCase)

	Routers(server, controller)
	server.Run(":8000")
}
