// Ponto de entrada do sistema responsável por tratar as requests e responses (dependem de UseCases)
package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"app/model"
	"app/pkg"
	"app/usecase"
)

type ToDo_Controller struct {
	ToDoUseCase usecase.ToDo_UseCase
}

// -- Constructor
func Init(usecase usecase.ToDo_UseCase) *ToDo_Controller {
	return &ToDo_Controller{
		ToDoUseCase: usecase,
	}
}

// -- Methods
func (it *ToDo_Controller) Create(ctx *gin.Context) {
	// tratando `request`
	var request model.ToDo
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	// Executando useCase
	err := it.ToDoUseCase.Create_ToDo(request)
	if err != nil {
		fmt.Println("Erro useCase")
		ctx.JSON(http.StatusBadRequest, "Erro!")
		return
	}

	// emitindo `response` 200 (StatusOK)
	ctx.JSON(http.StatusOK, "ToDo Criado")
}

func (it *ToDo_Controller) Read(ctx *gin.Context) {
	// tratando `request`
	request := ctx.Param("id")
	if request == "" {
		ctx.JSON(http.StatusBadRequest, "Id não fornecido")
	}

	idParam, err := strconv.Atoi(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id inválido")
	}

	// Executando useCase
	response, err := it.ToDoUseCase.Read_ToDo(idParam)
	if err != nil {
		fmt.Println("Erro useCase")
		ctx.JSON(http.StatusBadRequest, "Erro!")
		return
	}

	// emitindo `response` 200 (StatusOK)
	ctx.JSON(http.StatusOK, response)
}

func (it *ToDo_Controller) Read_All(ctx *gin.Context) {
	// Validação JWT
	var _token = ctx.Request.Header.Get("Authorization")
	if !pkg.ValidateJWT(_token) {
		ctx.JSON(403, "Token não autorizado")
		return
	}

	// Executando useCase
	var response, err = it.ToDoUseCase.Read_ToDoAll()
	if err != nil {
		fmt.Println("Erro UseCase", err)
		ctx.JSON(http.StatusBadRequest, "Erro InternalServer")
		return
	}

	// emitindo `response` 200 (StatusOK)
	ctx.JSON(http.StatusOK, response)
}

func (it *ToDo_Controller) Update(ctx *gin.Context) {
	// tratando `request`
	var request model.ToDo
	if err := ctx.BindJSON(&request); err != nil {
		fmt.Println("Erro na leitura da requisição")
		ctx.JSON(http.StatusBadRequest, "Erro na leitura da requisição")
		return
	}

	// executando useCase
	err := it.ToDoUseCase.Update_ToDo(request)
	if err != nil {
		fmt.Println("Erro useCase")
		ctx.JSON(http.StatusBadRequest, "Erro")
		return
	}

	// emitindo `response`
	ctx.JSON(http.StatusOK, "ToDo atualizado")
}

func (it *ToDo_Controller) Delete(ctx *gin.Context) {
	// tratando `request`
	request := ctx.Param("id")
	if request == "" {
		ctx.JSON(http.StatusBadRequest, "Id não fornecido")
	}

	idParam, err := strconv.Atoi(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "Id inválido")
	}

	// Executando useCase
	erro := it.ToDoUseCase.Delete_ToDo(idParam)
	if erro != nil {
		fmt.Println("Erro useCase")
		ctx.JSON(http.StatusBadRequest, "Erro!")
		return
	}

	// emitindo `response` 200 (StatusOK)
	ctx.JSON(http.StatusOK, "ToDo deletado")
}
