package models

import "gorm.io/gorm"

type Cliente struct {
    gorm.Model
    Nome    string `json:"nome"`
    Email   string `json:"email" gorm:"unique"`
    CPF     string `json:"cpf" gorm:"unique"`
    CEP     string `json:"cep"`
    Rua     string `json:"rua"`
    Numero  string `json:"numero"`
    Bairro  string `json:"bairro"`
    Cidade  string `json:"cidade"`
    Estado  string `json:"estado"`
}
