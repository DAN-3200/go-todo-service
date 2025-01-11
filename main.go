package main

import (
	"app/controller"
	"app/db"
	"app/repository"
	"app/usecase"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Executando o servidor... \n")
	RunServerAPI()
}

func RunServerAPI() {
	// Conexão com banco de dados
	var Conn, err = db.ConnDB()
	defer Conn.Close()
	if err != nil {
		log.Fatal("Erro", err)
		return
	}

	// `Arquitetura Limpa` em Camadas
	var useControllers = controller.Init(
		usecase.Init(
			repository.Init(Conn),
		),
	)

	// Framework Gin - http requests
	var server = gin.Default()
	var toDo = server.Group("/ToDo")
	{
		toDo.POST("/", useControllers.Create)
		toDo.GET("/:id", useControllers.Read)
		toDo.GET("/", useControllers.Read_All)
		toDo.PUT("/", useControllers.Update)
		toDo.DELETE("/:id", useControllers.Delete)
	}
	server.Run(":8000")
}