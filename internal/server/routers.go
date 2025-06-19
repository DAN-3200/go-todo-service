package server

import (
	"github.com/gin-gonic/gin"

	"app/internal/controller"
)

func Routers(server *gin.Engine, handles *controller.LayerController) {
	toDo := server.Group("/todo")
	{
		toDo.POST("", handles.SaveToDo)
		toDo.GET("/:id", handles.GetToDo)
		toDo.GET("", handles.GetToDoList)
		toDo.PATCH("/:id", handles.EditToDo)
		toDo.DELETE("/:id", handles.DeleteToDo)
	}
}
