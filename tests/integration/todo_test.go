package integration_test

import (
	"app/internal/inner/entity"
	"app/internal/outer/persistence/repository"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

var uri = "postgres://test:test@localhost:6543/test?sslmode=disable"

//  End-to-End ~= "de ponta a ponta"
func Test_ToDoRepository_E2E(t *testing.T) {
	conn, err := sql.Open("postgres", uri)
	require.NoError(t, err, err)

	repo, err := repository.InitLayer(conn)
	require.NoError(t, err, err)

	err = repo.CreateTable()
	require.NoError(t, err, err)
	{
		// -- Save
		id, err := repo.Save(entity.ToDo{
			Title:   "new",
			Content: "new",
			Status:  false,
		})
		require.NoError(t, err)

		// -- Edit
		// title := "UP"
		// content := "UP"
		// Status := false

		// err = repo.Edit(id, entity.ToDo{
		// 	Title:   &title,
		// 	Content: &content,
		// 	Status:  &Status,
		// })
		// require.NoError(t, err)

		// -- Get
		result, err := repo.Get(id)
		require.NoError(t, err)
		t.Logf("\n ToDo obtido: %+v", result)

		// -- Delete
		err = repo.Delete(id)
		require.NoError(t, err)

		// tem que dar erro
		_, err = repo.Get(id)
		require.Error(t, err)
	}
}
