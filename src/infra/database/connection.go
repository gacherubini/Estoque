package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// Obter as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Montar a string de conexão
	connectionStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable", dbHost, dbUser, dbPassword, dbName)

	// Abrir a conexão com o banco de dados
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar no banco de dados: %v", err)
	}

	// Testar a conexão com o banco de dados
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("DB Ping Failed: %v", err)
	}

	log.Println("DB Connection started successfully")
	return db, nil
}
