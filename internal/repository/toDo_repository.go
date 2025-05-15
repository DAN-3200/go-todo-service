// Abstração de acesso aos dados (dependem de frameworks/bancos, usados por UseCases)
package repository

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importar drive

	"app/internal/model"
)

type ToDo_Repository struct {
	ConnectDB *sql.DB
}

// -- Contructor
func Init(connection *sql.DB) ToDo_Repository {
	return ToDo_Repository{
		ConnectDB: connection,
	}
}

// -- Methods
func (it *ToDo_Repository) Insert_ToDo_DB(newToDo model.ToDo) error {
	// Conexão com o Banco de dados
	Conn := it.ConnectDB

	query := `INSERT INTO ToDo (title, content, status) VALUES ($1, $2, $3)`
	_, err := Conn.Exec(query, newToDo.Title, newToDo.Content, newToDo.Status)
	if err != nil {
		fmt.Println(`Erro Insert! \n`, err)
		return err
	}

	return nil
}

func (it *ToDo_Repository) Select_ToDo_DB(id int) (model.ToDo, error) {
	// Conexão com o banco de dados
	Conn := it.ConnectDB

	// Consulta ao banco a partir dos parametros
	var toDo model.ToDo
	row := Conn.QueryRow(
		`SELECT id, title, content, status FROM ToDo WHERE id=$1`, id,
	)
	err := row.Scan(&toDo.Id, &toDo.Title, &toDo.Content, &toDo.Status)

	// Tratamento de erro de consulta
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Nenhum registro encontro.")
		} else {
			fmt.Println("Erro de consulta: ", err)
		}
		return model.ToDo{}, err
	}

	return toDo, nil
}

func (it *ToDo_Repository) Select_All_ToDo_DB() ([]model.ToDo, error) {
	var Conn = it.ConnectDB

	// Consulta
	var query = `SELECT id, title, content, status FROM ToDo`
	var rows, err = Conn.Query(query)
	if err != nil {
		fmt.Println("Erro de consulta | ", err)
		return []model.ToDo{}, err
	}

	var toDoList []model.ToDo
	var item model.ToDo

	for rows.Next() {
		var err = rows.Scan(
			&item.Id,
			&item.Title,
			&item.Content,
			&item.Status,
		)
		if err != nil {
			fmt.Println(err)
			return []model.ToDo{}, err
		}
		toDoList = append(toDoList, item)
	}
	rows.Close()

	return toDoList, nil
}

func (it *ToDo_Repository) Update_ToDo_DB(mToDo model.ToDo) error {
	Conn := it.ConnectDB

	query := `UPDATE ToDo SET title=$1, content=$2, status = $3 WHERE id = $4`
	_, err := Conn.Exec(query, mToDo.Title, mToDo.Content, mToDo.Status, mToDo.Id)
	if err != nil {
		fmt.Println("Erro ao atulizar campo da table: ", err)
		return err
	}

	return nil
}

func (it *ToDo_Repository) Delete_ToDo_DB(id int) error {
	Conn := it.ConnectDB

	query := `DELETE FROM ToDo WHERE id = $1`
	_, err := Conn.Exec(query, id)
	if err != nil {
		fmt.Println("Erro ao excluir: ", err)
		return err
	}

	return nil
}
