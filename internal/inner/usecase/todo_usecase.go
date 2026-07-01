package usecase

import (
	"fmt"

	"app/internal/inner/dto"
	"app/internal/inner/entity"
	"app/internal/inner/ports"
)

type LayerUseCase struct {
	TodoRepo ports.Irepository[entity.ToDo]
}

func InitLayer(repository ports.Irepository[entity.ToDo]) *LayerUseCase {
	return &LayerUseCase{
		TodoRepo: repository,
	}
}

// ------------------------------------------------------------------

func (it *LayerUseCase) SaveToDo(info *dto.ToDoReq) (int64, error) {
	todo, err := entity.NewToDo(info.Title, info.Content)
	if err != nil {
		return 0, err
	}

	id, err := it.TodoRepo.Save(todo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (it *LayerUseCase) GetToDo(id int64) (*dto.ToDoRes, error) {
	if id <= 0 {
		return nil, fmt.Errorf("id inválido")
	}

	result, err := it.TodoRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return dto.ToToDoRes(*result), err
}

func (it *LayerUseCase) GetToDoList() ([]dto.ToDoRes, error) {
	result, err := it.TodoRepo.GetList()
	if err != nil {
		return []dto.ToDoRes{}, err
	}

	return dto.ToToDoResList(*result), nil
}

func (it *LayerUseCase) EditToDo(id int64, info dto.ToDoEditReq) error {
	if id <= 0 {
		return fmt.Errorf("id inválido")
	}

	todo, err := it.TodoRepo.Get(id)
	if err != nil {
		return fmt.Errorf("não há elemento com tal id")
	}

	if info.Title != nil {
		todo.Title = *info.Title
	}

	if info.Content != nil {
		todo.Content = *info.Content
	}

	if info.Status != nil {
		todo.Status = *info.Status
	}

	err = it.TodoRepo.Edit(todo)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) DeleteToDo(id int64) error {
	if id <= 0 {
		return fmt.Errorf("id inválido")
	}

	err := it.TodoRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
