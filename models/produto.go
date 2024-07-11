package models

import (
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Descricao         string  `json:"descricao"`
	QuantidadeEstoque int     `json:"quantidade_estoque"`
	Valor             float64 `json:"valor"`
	CategoriaId       uint    `json:"categoria_id"`
	ProdutoImagem     string  `json:"produto-imagem"`
}
