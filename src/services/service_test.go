package services

import (
	"go-poc/models"
	"testing"
)

func TestGetUser_Default(t *testing.T) {
	user := GetUser()
	if user.ID != 1 {
		t.Errorf("esperado ID=1, obtido %d", user.ID)
	}
	if user.Name != "Exemplo" {
		t.Errorf("esperado Name='Exemplo', obtido '%s'", user.Name)
	}
}

// Exemplo de teste para outros cenários (caso GetUser evolua para aceitar parâmetros)
func TestGetUser_Empty(t *testing.T) {
	emptyUser := models.User{}
	if emptyUser.ID != 0 {
		t.Errorf("esperado ID=0, obtido %d", emptyUser.ID)
	}
	if emptyUser.Name != "" {
		t.Errorf("esperado Name='', obtido '%s'", emptyUser.Name)
	}
}
