package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importar drive
)

func Conn_Postgres() *sql.DB {
	// Conectar ao banco Postgres
	var Conn, err = sql.Open("postgres", os.Getenv("url_db"))
	if err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	// Verificar se há resposta a conexão
	if err := Conn.Ping(); err != nil {
		fmt.Printf("Erro: %v", err)
		return nil
	}

	return Conn
}