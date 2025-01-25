package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // Importar drive
)

func Conn_Postgres() (*sql.DB, error) {
	// Conectar ao banco Postgres
	var useDB, err_db = sql.Open("postgres", os.Getenv("url_db"))
	if err_db != nil {
		panic(err_db)
	}

	// Verificar se há resposta a conexão
	var err error = useDB.Ping()

	return useDB, err
}
