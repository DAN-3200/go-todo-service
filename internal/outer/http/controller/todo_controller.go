package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app/internal/inner/contracts"
	"app/internal/inner/usecase"
	"app/pkg/utils"
)

type LayerController struct {
	UseCase *usecase.LayerUseCase
}

// -- Constructor
func InitLayer(usecase *usecase.LayerUseCase) *LayerController {
	return &LayerController{
		UseCase: usecase,
	}
}

// ------------------------------------------------------------------

func (it *LayerController) SaveToDo(ctx *gin.Context) {
	request, err := utils.MapReqJSON[contracts.ToDoReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := it.UseCase.SaveToDo(*request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"msg": "ToDo Criado", "id": id})
}

func (it *LayerController) GetToDo(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID == "" {
		ctx.String(http.StatusBadRequest, "Id não fornecido")
		return
	}

	idParam, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	response, err := it.UseCase.GetToDo(int64(idParam))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *LayerController) GetToDoList(ctx *gin.Context) {
	response, err := it.UseCase.GetToDoList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (it *LayerController) EditToDo(ctx *gin.Context) {
	paramID := ctx.Param("id")
	if paramID == "" {
		ctx.String(http.StatusBadRequest, "Id não fornecido")
		return
	}

	request, err := utils.MapReqJSON[contracts.ToDoEditReq](ctx)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	id, err := strconv.Atoi(paramID)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.UseCase.EditToDo(int64(id), *request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "ToDo atualizado")
}

func (it *LayerController) DeleteToDo(ctx *gin.Context) {
	request := ctx.Param("id")
	if request == "" {
		ctx.String(http.StatusBadRequest, "Id não fornecido")
		return
	}

	idParam, err := strconv.Atoi(request)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = it.UseCase.DeleteToDo(int64(idParam))
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.String(http.StatusOK, "ToDo deletado")
}
