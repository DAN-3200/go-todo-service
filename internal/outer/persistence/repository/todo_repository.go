package repository

import (
	"database/sql"
	"errors"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	// Importar drive
	_ "github.com/lib/pq"

	"app/internal/inner/entity"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

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

func (it *LayerRepository) Save(info *entity.ToDo) (int64, error) {
	q := psql.Insert("todo").Columns("title", "content", "status").Values(info.Title, info.Content, info.Status).Suffix("RETURNING id")
	query, args, err := q.ToSql()

	var id int64
	err = it.DB.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (it *LayerRepository) Get(id int64) (*entity.ToDo, error) {
	q := psql.Select("id", "title", "content", "status", "created_at").From("todo").Where(sq.Eq{"id": id})
	query, args, err := q.ToSql()
	row := it.DB.QueryRow(query, args...)

	var todo entity.ToDo
	err = row.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Content,
		&todo.Status,
		&todo.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Nenhum registro encontro.")
		}
		return nil, err
	}

	return &todo, nil
}

func (it *LayerRepository) GetList() (*[]entity.ToDo, error) {
	q := psql.Select("id", "title", "content", "status", "created_at").From("todo")
	query, _, err := q.ToSql()
	rows, err := it.DB.Query(query)
	if rows == nil {
		return nil, fmt.Errorf("Nenhum linha afetada")
	}
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var todoList []entity.ToDo
	var todo entity.ToDo
	for rows.Next() {
		var err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.Status,
			&todo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		todoList = append(todoList, todo)
	}

	return &todoList, nil
}

func (it *LayerRepository) Edit(info *entity.ToDo) error {
	q := psql.Update("todo").
		Set("title", info.Title).
		Set("content", info.Content).
		Set("status", info.Status).
		Where(sq.Eq{"id": info.ID})

	sql, args, err := q.ToSql()
	if err != nil {
		return err
	}

	_, err = it.DB.Exec(sql, args...)
	if err != nil {
		return err
	}

	return nil
}

func (it *LayerRepository) Delete(id int64) error {
	q := psql.Delete("todo").Where(sq.Eq{"id": id})

	sql, args, err := q.ToSql()
	if err != nil {
		return err
	}

	result, err := it.DB.Exec(sql, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Nenhum registro afetado")
	}

	return nil
}

func (it *LayerRepository) CreateTable() error {
	_, err := it.DB.Exec(`
		CREATE TABLE IF NOT EXISTS todo (
			id SERIAL PRIMARY KEY,
			title TEXT,
			content TEXT,
			status BOOLEAN,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);	
	`,
	)

	if err != nil {
		return err
	}
	return nil
}
