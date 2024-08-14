package routes

import (
	"neoway/backend/controllers"
	"neoway/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes configura as rotas da aplicação
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// Rota raiz que fornece informações sobre a aplicação
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bem-vindo à API Neoway! Esta API oferece endpoints para gerenciar clientes e consultar o status do sistema.",
			"endpoints": []string{
				"/clients - Gerencia clientes",
				"/status - Verifica o status da API",
			},
		})
	})

	// Configuração das rotas para operações com clientes
	clientRepo := &models.GORMClientRepository{DB: db}
	clientController := controllers.NewClientController(clientRepo)
	clientController.RegisterRoutes(r)

	// Configuração da rota para obter status do servidor
	statusController := controllers.NewStatusController()
	statusController.RegisterRoutes(r)
}
