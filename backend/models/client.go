package models

import (
	"errors"
	"time"
)

// Client representa o modelo de cliente
type Client struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	CPF_CNPJ  string    `gorm:"unique;not null" json:"cpf_cnpj"`
	Name      string    `gorm:"not null" json:"name"`
	IsBlocked bool      `json:"is_blocked"`
	CreatedAt time.Time `json:"created_at"`
}

// ErrRecordNotFound é retornado quando um registro não é encontrado.
var ErrRecordNotFound = errors.New("record not found")
