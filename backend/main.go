package main

import (
	"neoway/backend/database"
	"neoway/backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa o banco de dados
	db, err := database.Connect()
	if err != nil {
		panic("failed to connect database")
	}

	// Cria o roteador do Gin
	r := gin.Default()

	// Habilita o CORS
	r.Use(cors.Default())

	// Configura as rotas
	routes.SetupRoutes(r, db)

	// Inicia o servidor
	r.Run(":8080")
}
