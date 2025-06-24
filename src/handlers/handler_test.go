package handlers

import (
	"encoding/json"
	"go-poc/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/user", nil)
	rr := httptest.NewRecorder()

	UserHandler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("esperado status %d, obtido %d", http.StatusOK, status)
	}

	var user models.User
	err := json.NewDecoder(rr.Body).Decode(&user)
	if err != nil {
		t.Fatalf("erro ao decodificar resposta: %v", err)
	}

	if user.ID != 1 || user.Name != "Exemplo" {
		t.Errorf("esperado user {ID:1, Name:'Exemplo'}, obtido %+v", user)
	}
}
