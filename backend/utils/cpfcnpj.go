package utils

import (
	"errors"
	"strings"

	"github.com/klassmann/cpfcnpj"
)

// ValidateDocument valida o CPF ou CNPJ
func ValidateDocument(document string) (bool, error) {
	document = strings.TrimSpace(document)
	document = cpfcnpj.Clean(document) // Remove caracteres não numéricos

	if len(document) == 11 {
		// CPF
		return ValidateCPF(document) == nil, nil
	} else if len(document) == 14 {
		// CNPJ
		return ValidateCNPJ(document) == nil, nil
	}
	return false, errors.New("número de dígitos inválido")
}

// ValidateCPF verifica se o CPF é válido
func ValidateCPF(cpf string) error {
	if !cpfcnpj.ValidateCPF(cpf) {
		return errors.New("CPF inválido")
	}
	return nil
}

// ValidateCNPJ verifica se o CNPJ é válido
func ValidateCNPJ(cnpj string) error {
	if !cpfcnpj.ValidateCNPJ(cnpj) {
		return errors.New("CNPJ inválido")
	}
	return nil
}

// NormalizeCPF formata o CPF no formato padrão (XXX.XXX.XXX-XX)
func NormalizeCPF(cpf string) (string, error) {
	cpf = strings.TrimSpace(cpf)
	cpf = cpfcnpj.Clean(cpf) // Remove caracteres não numéricos
	if !cpfcnpj.ValidateCPF(cpf) {
		return "", errors.New("CPF inválido")
	}
	return formatCPF(cpf), nil
}

// NormalizeCNPJ formata o CNPJ no formato padrão (XX.XXX.XXX/0001-XX)
func NormalizeCNPJ(cnpj string) (string, error) {
	cnpj = strings.TrimSpace(cnpj)
	cnpj = cpfcnpj.Clean(cnpj) // Remove caracteres não numéricos
	if !cpfcnpj.ValidateCNPJ(cnpj) {
		return "", errors.New("CNPJ inválido")
	}
	return formatCNPJ(cnpj), nil
}

// formatCPF formata o CPF no formato padrão (XXX.XXX.XXX-XX)
func formatCPF(cpf string) string {
	if len(cpf) != 11 {
		return cpf // Ou você pode retornar um erro, dependendo do caso de uso
	}
	return cpf[:3] + "." + cpf[3:6] + "." + cpf[6:9] + "-" + cpf[9:]
}

// formatCNPJ formata o CNPJ no formato padrão (XX.XXX.XXX/0001-XX)
func formatCNPJ(cnpj string) string {
	if len(cnpj) != 14 {
		return cnpj // Ou você pode retornar um erro, dependendo do caso de uso
	}
	return cnpj[:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:]
}
