package handlers

import (
	"api-pdv-golang/config"
	"api-pdv-golang/models"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	var usuarios []models.Usuario
	config.DB.Find(&usuarios)
	json.NewEncoder(w).Encode(usuarios) 
}

func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost )
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	usuario.Senha = string(hashedPassword)
	
	config.DB.Create(&usuario)
	json.NewEncoder(w).Encode(usuario)
}