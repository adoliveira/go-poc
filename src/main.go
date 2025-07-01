package main

import (
	"fmt"
	"go-poc/handlers"
	"net/http"
)

// StartServer inicializa e executa o servidor HTTP da API
func StartServer(mux http.Handler) {
	fmt.Println("API Go rodando na porta 8080...")
	http.ListenAndServe(":8080", mux)
}

// BuilServerMux cria e retorna um *http.ServeMux com as rotas da API
func BuilServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", handlers.UserHandler)
	return mux
}

func main() {
	StartServer(BuilServerMux())
}
