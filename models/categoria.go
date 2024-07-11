package models

import "gorm.io/gorm"


type Categoria struct {
	gorm.Model
	Descricao string `json:"descricao"`
}