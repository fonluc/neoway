package controllers

import (
	"fmt"
	"neoway/backend/models"
	"neoway/backend/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/klassmann/cpfcnpj"
)

type ClientController struct {
	repo models.ClientRepository
}

// NewClientController cria uma nova instância do ClientController
func NewClientController(repo models.ClientRepository) *ClientController {
	return &ClientController{repo: repo}
}

// RegisterRoutes registra as rotas no gin router
func (c *ClientController) RegisterRoutes(r *gin.Engine) {
	r.POST("/clients", c.CreateClient)
	r.GET("/clients", c.ListClients)
	r.GET("/clients/:cpfcnpj", c.GetClientByCPF)
	r.GET("/clients/search", c.SearchClients)
}

// CreateClient adiciona um novo cliente ao banco de dados
func (c *ClientController) CreateClient(ctx *gin.Context) {
	var client models.Client
	if err := ctx.ShouldBindJSON(&client); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Limpa e normaliza o CPF/CNPJ
	document := strings.TrimSpace(client.CPF_CNPJ)
	document = cpfcnpj.Clean(document)
	client.CPF_CNPJ = document

	// Valida o CPF/CNPJ
	valid, err := utils.ValidateDocument(document)
	if !valid || err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "CPF/CNPJ inválido"})
		return
	}

	// Verifica se o cliente já está cadastrado
	_, err = c.repo.GetClientByCPF(document)
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Cliente já cadastrado"})
		return
	} else if err.Error() != "record not found" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao verificar cliente"})
		return
	}

	// Cria o cliente
	if err := c.repo.CreateClient(client); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar cliente"})
		return
	}

	ctx.JSON(http.StatusCreated, client)
}

// ListClients retorna todos os clientes cadastrados
func (c *ClientController) ListClients(ctx *gin.Context) {
	clients, err := c.repo.ListClients()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar clientes"})
		return
	}

	// Log the clients to verify
	fmt.Println("Clients:", clients)

	ctx.JSON(http.StatusOK, clients)
}

// GetClientByCPF retorna um cliente específico baseado no CPF/CNPJ
func (c *ClientController) GetClientByCPF(ctx *gin.Context) {
	cpfcnpjParam := ctx.Param("cpfcnpj")
	cpfcnpjParam = cpfcnpj.Clean(cpfcnpjParam) // Limpa o CPF/CNPJ recebido
	client, err := c.repo.GetClientByCPF(cpfcnpjParam)
	if err != nil {
		if err.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar cliente"})
		}
		return
	}
	ctx.JSON(http.StatusOK, client)
}

// SearchClients busca clientes pelo nome/razão social e ordena alfabeticamente
func (c *ClientController) SearchClients(ctx *gin.Context) {
	name := ctx.Query("name")
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Nome não fornecido"})
		return
	}

	clients, err := c.repo.ListClients()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar clientes"})
		return
	}

	var filteredClients []models.Client
	for _, client := range clients {
		if containsIgnoreCase(client.Name, name) {
			filteredClients = append(filteredClients, client)
		}
	}

	ctx.JSON(http.StatusOK, filteredClients)
}

// Helper function to check if a string contains another string, ignoring case
func containsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
