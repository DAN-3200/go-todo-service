package db

import (
	"fmt"
	"os"

	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Importar drive
)

func ConnDB() (*sql.DB, error) {
	// Carregar variáveis de ambiente
	err_load := godotenv.Load()
	if err_load != nil {
		fmt.Println("Erro no godotenv.Load")
		panic(err_load)
	}

	// Conectar ao banco Postgres
	DB, err_db := sql.Open("postgres", os.Getenv("url_db"))
	if err_db != nil {
		panic(err_db)
	}

	// Verificar se há resposta a conexão
	var err error = DB.Ping()

	return DB, err
}