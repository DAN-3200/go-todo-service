package contracts

import (
	"app/internal/dto"
)

type RepoSQL interface {
	SaveToDo(newToDo dto.ToDoReq) (int64, error)
	GetToDo(id int64) (dto.ToDoRes, error)
	GetToDoList() ([]dto.ToDoRes, error)
	EditToDo(id int64, newInfo dto.ToDoEditReq) error
	DeleteToDo(id int64) error
}
