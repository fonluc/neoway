package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ClientRepository define os métodos necessários para manipular clientes
type ClientRepository interface {
	CreateClient(client Client) error
	GetClientByCPF(cpfcnpj string) (*Client, error)
	ListClients() ([]Client, error)
}

// NewGORMClientRepository cria uma nova instância de GORMClientRepository
func NewGORMClientRepository(db *gorm.DB) *GORMClientRepository {
	return &GORMClientRepository{DB: db}
}

// GORMClientRepository implementa a interface ClientRepository usando GORM
type GORMClientRepository struct {
	DB *gorm.DB
}

// / NewClientRepository cria uma nova instância do GORMClientRepository
func NewClientRepository() *GORMClientRepository {
	dsn := "user=postgres password=@postgresql dbname=neoway_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Migra o esquema, se necessário
	if err := db.AutoMigrate(&Client{}); err != nil {
		panic("failed to migrate database schema: " + err.Error())
	}

	return &GORMClientRepository{DB: db}
}

// CreateClient cria um novo cliente no banco de dados
func (r *GORMClientRepository) CreateClient(client Client) error {
	return r.DB.Create(&client).Error
}

// GetClientByCPF busca um cliente pelo CPF/CNPJ
func (r *GORMClientRepository) GetClientByCPF(cpfcnpj string) (*Client, error) {
	var client Client
	if err := r.DB.Where("cpf_cnpj = ?", cpfcnpj).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

// ListClients lista todos os clientes
func (r *GORMClientRepository) ListClients() ([]Client, error) {
	var clients []Client
	if err := r.DB.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}
