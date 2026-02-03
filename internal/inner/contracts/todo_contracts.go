package contracts

import (
	"app/internal/inner/entity"
	"time"
)

type ToDoReq struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	Status  bool   `json:"status"`
}

type ToDoRes struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type ToDoEditReq struct {
	Title   *string `json:"title"`
	Content *string `json:"content"`
	Status  *bool   `json:"status"`
}

func ToToDoRes(t entity.ToDo) ToDoRes {
	return ToDoRes{
		ID:        t.ID,
		Title:     t.Title,
		Content:   t.Content,
		Status:    t.Status,
		CreatedAt: t.CreatedAt,
	}
}

func ToToDoResList(list []entity.ToDo) []ToDoRes {
	res := make([]ToDoRes, 0, len(list))

	for _, t := range list {
		res = append(res, ToToDoRes(t))
	}

	return res
}
