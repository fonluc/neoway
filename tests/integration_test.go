package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Função para criar uma conexão com o banco de dados de teste
func NewTestDB() (*gorm.DB, error) {
	// DSN para o banco de dados de teste
	dsn := "host=localhost user=postgres password=@postgresql dbname=neoway_test sslmode=disable"

	// Abrir a conexão usando o driver PostgreSQL do GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Configuração opcional de log
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm DB: %w", err)
	}

	// Verificar se a conexão está funcionando
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB from gorm DB: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil
}

// Função para configurar o roteador
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	// Adiciona uma rota simples para testar o servidor
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	return r
}

// Função de teste de integração
func TestIntegration(t *testing.T) {
	// Configuração do banco de dados de teste
	db, err := NewTestDB()
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Exec("DROP DATABASE IF EXISTS neoway_test")
	}()

	// Teste para verificar a conexão com o banco de dados
	t.Run("Database Connection", func(t *testing.T) {
		sqlDB, err := db.DB()
		if err != nil {
			t.Fatalf("failed to get sql.DB from gorm DB: %v", err)
		}

		// Testar uma consulta simples para verificar a conexão
		var count int
		err = sqlDB.QueryRow("SELECT COUNT(*) FROM pg_catalog.pg_tables").Scan(&count)
		if err != nil {
			t.Fatalf("failed to query database: %v", err)
		}

		assert.Greater(t, count, 0, "Database should have at least one table")
	})

	// Teste para verificar o servidor
	router := setupRouter()

	t.Run("Server Ping", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/ping", nil)
		recorder := httptest.NewRecorder()
		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.JSONEq(t, `{"message":"pong"}`, recorder.Body.String())
	})
}
