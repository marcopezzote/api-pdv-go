package models

import "gorm.io/gorm"

type PedidoProduto struct {
    gorm.Model
    PedidoID  uint `json:"pedido_id"`
    ProdutoID uint `json:"produto_id"`
}
