package handlers

import (
	"api-pdv-golang/config"
	"api-pdv-golang/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)


type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
    var creds struct {
        Email string `json:"email"`
        Senha string `json:"senha"`
    }

    json.NewDecoder(r.Body).Decode(&creds)

    var usuario models.Usuario
    if err := config.DB.Where("email = ?", creds.Email).First(&usuario).Error; err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(usuario.Senha), []byte(creds.Senha)); err != nil {
        http.Error(w, "Invalid password", http.StatusUnauthorized)
        return
    }

    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: creds.Email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(config.JwtKey)
    if err != nil {
        http.Error(w, "Could not create token", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name:    "token",
        Value:   tokenString,
        Expires: expirationTime,
    })

    json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}