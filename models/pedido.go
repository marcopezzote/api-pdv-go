package models

import "gorm.io/gorm"

type Pedido struct {
    gorm.Model
    ClienteID  uint    `json:"cliente_id"`
    Observacao string  `json:"observacao"`
    ValorTotal float64 `json:"valor_total"`
}
