package usecase

import (
	"app/internal/contracts"
	"app/internal/dto"
)

type LayerUseCase struct {
	Repo contracts.RepoSQL
}

func InitLayer(repository contracts.RepoSQL) *LayerUseCase {
	return &LayerUseCase{
		Repo: repository,
	}
}

// ------------------------------------------------------------------

func (it *LayerUseCase) SaveToDo(newTodo dto.ToDoReq) (int64, error) {
	id, err := it.Repo.SaveToDo(newTodo)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (it *LayerUseCase) GetToDo(id int64) (dto.ToDoRes, error) {
	result, err := it.Repo.GetToDo(id)
	if err != nil {
		return dto.ToDoRes{}, err
	}

	return result, err
}

func (it *LayerUseCase) GetToDoList() ([]dto.ToDoRes, error) {
	result, err := it.Repo.GetToDoList()
	if err != nil {
		return []dto.ToDoRes{}, err
	}
	return result, nil
}

func (it *LayerUseCase) EditToDo(id int64, newInfo dto.ToDoEditReq) error {
	err := it.Repo.EditToDo(id, newInfo)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerUseCase) DeleteToDo(id int64) error {
	err := it.Repo.DeleteToDo(id)
	if err != nil {
		return err
	}

	return nil
}
