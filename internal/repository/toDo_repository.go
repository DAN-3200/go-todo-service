package repository

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq" // Importar drive

	"app/internal/dto"
	"app/pkg/utils"
)

type LayerRepository struct {
	DB *sql.DB
}

func InitLayer(connection *sql.DB) (*LayerRepository, error) {
	if err := connection.Ping(); err != nil {
		return &LayerRepository{}, err
	}
	return &LayerRepository{
		DB: connection,
	}, nil
}

// ------------------------------------------------------------------

func (it *LayerRepository) SaveToDo(newToDo dto.ToDoReq) (int64, error) {
	query := `INSERT INTO ToDo (title, content, status) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	err := it.DB.QueryRow(query, newToDo.Title, newToDo.Content, newToDo.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (it *LayerRepository) GetToDo(id int64) (dto.ToDoRes, error) {
	row := it.DB.QueryRow(`SELECT id, title, content, status, created_at FROM ToDo WHERE id=$1`, id)

	var todo dto.ToDoRes
	err := row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.Status,
		&todo.CreatedAt,
	)

	// Tratamento de erro de consulta
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontro.")
		} else {
			fmt.Println("Erro de consulta: ", err)
		}
		return dto.ToDoRes{}, err
	}

	return todo, nil
}

func (it *LayerRepository) GetToDoList() ([]dto.ToDoRes, error) {
	query := `SELECT id, title, content, status, created_at FROM ToDo`
	rows, err := it.DB.Query(query)
	defer rows.Close()

	if err != nil {
		return []dto.ToDoRes{}, err
	}

	var todoList []dto.ToDoRes
	var todo dto.ToDoRes
	for rows.Next() {
		var err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.Status,
			&todo.CreatedAt,
		)
		if err != nil {
			return []dto.ToDoRes{}, err
		}
		todoList = append(todoList, todo)
	}

	return todoList, nil
}

func (it *LayerRepository) EditToDo(id int64, newInfo dto.ToDoEditReq) error {
	cols, args, err := utils.MapSQLInsertFields(
		map[string]string{
			"Title":   "title",
			"Content": "content",
			"Status":  "status",
		},
		newInfo,
	)
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`UPDATE ToDo SET %s WHERE id=%d`, strings.Join(cols, ", "), id)
	// fmt.Println(query)
	_, err = it.DB.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerRepository) DeleteToDo(id int64) error {
	query := `DELETE FROM ToDo WHERE id=$1`
	_, err := it.DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerRepository) CreateTable() error {
	_, err := it.DB.Exec(`
		CREATE TABLE IF NOT EXISTS ToDo (
			id SERIAL PRIMARY KEY,
			title TEXT,
			content TEXT,
			status BOOLEAN,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);	
	`,
	)

	if err != nil {
		fmt.Println("Erro: ", err)
		return err
	}
	return nil
}
