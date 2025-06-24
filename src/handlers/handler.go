package handlers

import (
	"encoding/json"
	"go-poc/services"
	"net/http"
)

// UserHandler responde com um usuário de exemplo
func UserHandler(w http.ResponseWriter, r *http.Request) {
	user := services.GetUser()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
