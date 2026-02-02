package entity

import (
	"time"
	"strings"
)


type ToDo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewToDo(title string, content string) *ToDo {
	return &ToDo{
		ID: 0,
		Title: strings.ToLower(title),
		Content: strings.ToLower(content),
		Status: false,
		CreatedAt: time.Now(),
	}
}


