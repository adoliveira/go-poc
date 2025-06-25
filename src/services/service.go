package services

import "go-poc/models"

// GetUser retorna um usuário de exemplo
func GetUser() models.User {
	return models.User{ID: 1, Name: "Exemplo"}
}
