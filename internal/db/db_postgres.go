package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importar drive
)

func ConnPostgreSQL() *sql.DB {
	// postgres://<usuário>:<senha>@<host>:<porta>/<banco_de_dados>?sslmode=<modo_ssl>
	uri := fmt.Sprintf(
		"postgres://%s:%s@localhost:5800/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	conn, err := sql.Open("postgres", uri)
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	if err := conn.Ping(); err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	return conn
}
