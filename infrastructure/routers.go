package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"app/controller"
	"app/pkg"
)

func Routers(server *gin.Engine, useControllers *controller.ToDo_Controller) {

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
