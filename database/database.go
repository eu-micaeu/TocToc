package database

// Importando bibliotecas para a criação e conexão do banco de dados
import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	_ "github.com/lib/pq"
)

// Função para criar e conectar a um banco de dados
func NewDB() (*sql.DB, error) {

	dbUser := "toctoc_user"
	dbPassword := "wWVNys6YIq4OzTqqs0dT5wMmUshzUGL0"
	dbHost := "dpg-cktkq4unfb1c73eh2h8g-a.oregon-postgres.render.com"
	dbPort := "5432"
	dbName := "toctoc"

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", dbUser, dbPassword, dbHost, dbPort, dbName) // String utilizada para criar a URL da conexão

	db, err := sql.Open("postgres", dsn) // Função SQL para abrir a conexão

	if err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Erro ao conectar ao banco de dados:", err)
		return nil, err
	}
	
	log.Println("Conexão com o banco de dados estabelecida com sucesso!")
	

	return db, nil
}


