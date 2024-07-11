package handlers

import (
	"api-pdv-golang/config"
	"api-pdv-golang/models"
	"encoding/json"
	"net/http"
	"strings"

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

	if strings.TrimSpace(usuario.Nome) == "" || strings.TrimSpace(usuario.Email) == "" || strings.TrimSpace(usuario.Senha) == "" {
		http.Error(w, "Nome, email e senha são obrigatórios!", http.StatusBadRequest)
		return
	}

	var existingUser models.Usuario
	if err := config.DB.Where("email = ? ", usuario.Email).First(&existingUser).Error; err == nil {
		http.Error(w, "Email already in use!", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usuario.Senha), bcrypt.DefaultCost )
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	usuario.Senha = string(hashedPassword)

	if err := config.DB.Create(&usuario).Error; err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}
	
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(usuario)
}