package main

import (
	"estoque/src/application/api"
	"estoque/src/infra/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	ginApp := gin.Default()

	ginApp.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Iniciar o servidor
	api.Server("3000", ginApp)
}
