package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func NewDB() (*sql.DB, error) {

	dbUser := "root"
	dbPassword := "wAAIT5t8pqS3YozQ1i698A7nstL6Kft2"
	dbHost := "dpg-clnrdu0fvntc739faddg-a.oregon-postgres.render.com"
	dbPort := "5432"
	dbName := "toctoc"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName) // String utilizada para criar a URL da conexão

	db, err := sql.Open("postgres", dsn) // Função SQL para abrir a conexão

	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	return db, nil
}
