package services

import "go-poc/models"

// UserService representa um serviço de exemplo
func GetUser() models.User {
	return models.User{ID: 1, Name: "Exemplo"}
}
