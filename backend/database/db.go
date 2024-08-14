package database

import (
	"fmt"
	"log"

	"neoway/backend/models"

	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Connect configura a conexão com o banco de dados PostgreSQL e retorna a instância
func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=@postgresql dbname=neoway_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados: %v", err)
	}

	// Criação de tabelas
	if err := createTables(db); err != nil {
		return nil, fmt.Errorf("erro ao criar tabelas: %v", err)
	}

	log.Println("Banco de dados conectado e tabelas criadas com sucesso.")
	return db, nil
}

// createTables cria as tabelas necessárias no banco de dados
func createTables(db *gorm.DB) error {
	// Usando AutoMigrate do GORM para criar tabelas com base nos modelos
	if err := db.AutoMigrate(&models.Client{}); err != nil {
		return fmt.Errorf("erro ao migrar tabelas: %v", err)
	}
	return nil
}

// CloseDB fecha a conexão com o banco de dados
func CloseDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("erro ao obter instância de SQL DB: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		return fmt.Errorf("erro ao fechar a conexão com o banco de dados: %v", err)
	}
	return nil
}

// NewTestDB cria e retorna uma instância de banco de dados PostgreSQL para testes
func NewTestDB() (*gorm.DB, error) {
	// Conecta ao banco de dados PostgreSQL padrão para criar o banco de dados de teste
	db, err := sql.Open("postgres", "host=localhost user=postgres password=@postgresql sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados PostgreSQL: %v", err)
	}
	defer db.Close()

	// Cria o banco de dados de teste se não existir
	_, err = db.Exec("CREATE DATABASE neoway_test")
	if err != nil && err.Error() != "pq: database \"neoway_test\" already exists" {
		return nil, fmt.Errorf("erro ao criar o banco de dados de teste: %v", err)
	}

	// Conecta ao banco de dados de teste
	dsn := "host=localhost user=postgres password=@postgresql dbname=neoway_test port=5432 sslmode=disable"
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir a conexão com o banco de dados de teste: %v", err)
	}

	// Criação de tabelas para o banco de dados de teste
	if err := createTables(gormDB); err != nil {
		return nil, fmt.Errorf("erro ao criar tabelas no banco de dados de teste: %v", err)
	}

	return gormDB, nil
}
