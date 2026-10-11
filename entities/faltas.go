package entities

import "time"

type Falta struct {
	ID int `json:"id" gorm:"primaryKey"`

	QuantidadeFalta int       `json:"quantidade_falta"`
	QuantidadeTotal int       `json:"quantidade_total"`
	Data            time.Time `json:"data"`

	ProdutoID int     `json:"produto_id"`
	Produto   Produto `json:"produto" gorm:"foreignKey:ProdutoID"`

	VendedorID int      `json:"vendedor_id"`
	Vendedor   Vendedor `json:"vendedor" gorm:"foreignKey:VendedorID"`
}