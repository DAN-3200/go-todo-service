package entity

import (
	"fmt"
	"strings"
	"time"
)

type ToDo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewToDo(title string, content string) (*ToDo, error) {
	obj := &ToDo{
		ID:        0,
		Title:     strings.ToLower(title),
		Content:   strings.ToLower(content),
		Status:    false,
		CreatedAt: time.Now(),
	}

	return obj, obj.Validate()
}

func (t *ToDo) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return fmt.Errorf("título é obrigatório")
	}
	if strings.TrimSpace(t.Content) == "" {
		return fmt.Errorf("conteúdo é obrigatório")
	}
	return nil
}
