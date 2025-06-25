package main

import (
	"fmt"
	"go-poc/handlers"
	"net/http"
)

// StartServer inicializa e executa o servidor HTTP da API
func StartServer() {
	fmt.Println("API Go rodando na porta 8080...")
	http.HandleFunc("/user", handlers.UserHandler)
	http.ListenAndServe(":8080", nil)
}

func main() {
	StartServer()
}
