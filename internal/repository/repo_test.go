package repository_test

import (
	"app/internal/dto"
	"app/internal/repository"
	"app/pkg/utils"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/mattn/go-sqlite3"

	"github.com/stretchr/testify/require"
)

var uri = "postgres://test:test@localhost:6543/test?sslmode=disable"

func Test_LayerRepositoryConnection(t *testing.T) {
	InitRepo(t)
}

func InitRepo(t *testing.T) *repository.LayerRepository {
	conn, err := sql.Open("postgres", uri)
	require.NoError(t, err, err)

	repo, err := repository.InitLayer(conn)
	require.NoError(t, err, err)

	err = repo.CreateTable()
	require.NoError(t, err, err)

	return repo
}

func Test_SaveToDo(t *testing.T) {
	repo := InitRepo(t)
	_, err := repo.SaveToDo(dto.ToDoReq{
		Title:   "Title",
		Content: "Content",
		Status:  false,
	})
	require.NoError(t, err)
}

func Test_GetToDo(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	repo, err := repository.InitLayer(db)
	require.NoError(t, err)

	// Dados simulados que o SELECT deve retornar
	rows := sqlmock.NewRows([]string{"id", "title", "content", "status", "created_at"}).
		AddRow(1, "newTitle", "newContent", false, time.Now())

	mock.ExpectQuery("SELECT id, title, content, status, created_at FROM ToDo WHERE id=\\$1").
		WithArgs(int64(1)).
		WillReturnRows(rows)

	result, err := repo.GetToDo(1)
	require.NoError(t, err)

	require.Equal(t, int64(1), result.ID)
	require.Equal(t, "newTitle", result.Title)
	require.Equal(t, "newContent", result.Content)
	require.False(t, result.Status)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func Test_GetToDoList(t *testing.T) {
	repo := InitRepo(t)

	list, err := repo.GetToDoList()
	require.NoError(t, err)

	t.Logf("\n ToDoList obtido: %+v", list)
}

func Test_EditToDo(t *testing.T) {
	repo := InitRepo(t)

	id, err := repo.SaveToDo(dto.ToDoReq{
		Title:   "new",
		Content: "new",
		Status:  false,
	})
	require.NoError(t, err)

	title := "UP"
	content := "UP"
	Status := false

	err = repo.EditToDo(id, dto.ToDoEditReq{
		Title:   &title,
		Content: &content,
		Status:  &Status,
	})
	require.NoError(t, err)

	result, err := repo.GetToDo(id)
	require.NoError(t, err)

	t.Logf("\n ToDo obtido: %+v", result)
}

func Test_DeleteToDo(t *testing.T) {
	repo := InitRepo(t)

	id, err := repo.SaveToDo(dto.ToDoReq{
		Title:   "new",
		Content: "new",
		Status:  false,
	})
	require.NoError(t, err)

	err = repo.DeleteToDo(id)
	require.NoError(t, err)

	result, err := repo.GetToDo(id)
	// tem que dar erro
	require.Error(t, err)

	t.Logf("\n ToDo obtido: %+v", result)
}

func Test_FieldPATCH(t *testing.T) {
	// -- campos da requisição
	content := "green"
	status := false

	todo := dto.ToDoEditReq{
		Title:   nil,
		Content: &content,
		Status:  &status,
	}

	// -----

	cols, args, err := utils.MapSQLInsertFields(
		map[string]string{
			"Title":   "title",
			"Content": "content",
			"Status":  "status",
		},
		todo,
	)
	
	require.NoError(t, err)
	t.Log("\n", cols, args)
}
