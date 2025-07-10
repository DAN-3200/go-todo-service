// DTO (Data Transfer Object)
package dto

import "time"

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
