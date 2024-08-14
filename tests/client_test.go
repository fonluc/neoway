package tests

import (
	"bytes"
	"encoding/json"
	"neoway/backend/controllers"
	"neoway/backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock de ClientRepository
type MockClientRepository struct {
	mock.Mock
}

func (m *MockClientRepository) CreateClient(client models.Client) error {
	args := m.Called(client)
	return args.Error(0)
}

func (m *MockClientRepository) GetClientByCPF(cpfcnpj string) (*models.Client, error) {
	args := m.Called(cpfcnpj)
	if client, ok := args.Get(0).(*models.Client); ok {
		return client, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockClientRepository) ListClients() ([]models.Client, error) {
	args := m.Called()
	return args.Get(0).([]models.Client), args.Error(1)
}

func TestCreateClient(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := new(MockClientRepository)
	clientController := controllers.NewClientController(mockRepo)
	router := gin.Default()
	clientController.RegisterRoutes(router)

	// Mock para retorno de cliente não encontrado
	mockRepo.On("GetClientByCPF", mock.Anything).Return(nil, models.ErrRecordNotFound)

	// Mock para sucesso na criação
	mockRepo.On("CreateClient", mock.Anything).Return(nil)

	clientData := models.Client{CPF_CNPJ: "74841780041", Name: "John Doe"}
	body, _ := json.Marshal(clientData)
	req, _ := http.NewRequest("POST", "/clients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
}
