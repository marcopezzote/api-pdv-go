package main

import (
	"api-pdv-golang/config"
	"api-pdv-golang/handlers"
	"api-pdv-golang/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
    // Carregar as variáveis de ambiente do arquivo .env
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.ConnectDatabase()

    r := mux.NewRouter()
    r.HandleFunc("/usuarios", handlers.GetUsuarios).Methods("GET")
    r.HandleFunc("/usuarios", handlers.CreateUsuario).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")

    // Rotas protegidas
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.JWTMiddleware)
    api.HandleFunc("/usuarios", handlers.GetUsuarios).Methods("GET")

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}