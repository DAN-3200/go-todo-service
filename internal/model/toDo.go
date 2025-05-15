// Núcleo fundamental do sistema (não depende de ninguém)
package model

type ToDo struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  bool   `json:"status"`
}