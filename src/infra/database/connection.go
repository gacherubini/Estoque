package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_NAME")

	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("Erro: uma ou mais variáveis de ambiente não estão configuradas corretamente")
	}

	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco de dados: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB Ping Failed: %v", err)
	}

	log.Println("DB Connection started successfully")
	return db, nil
}
