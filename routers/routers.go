package routers

import (
	"app/controller"

	"app/pkg"
	"app/repository"
	"app/usecase"
	"net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"database/sql"
)

func Init(server *gin.Engine, Conn *sql.DB) {
	// `Arquitetura Limpa` em Camadas
	var useControllers = controller.Init(
		usecase.Init(
			repository.Init(Conn),
		),
	)

	// Envia o JWT ao cliente
	server.GET("/key", func(ctx *gin.Context) {
		var tokenString, err = pkg.GenerateJWT()
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, "Erro interno do servidor")
			return
		}
		ctx.JSON(http.StatusOK, tokenString)
	})

	var toDo = server.Group("/ToDo")
	{
		toDo.POST("/", useControllers.Create)
		toDo.GET("/:id", useControllers.Read)
		toDo.GET("/", useControllers.Read_All)
		toDo.PUT("/", useControllers.Update)
		toDo.DELETE("/:id", useControllers.Delete)
	}
}
