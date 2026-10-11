package entities

type Produto struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Nome string `json:"nome"`
	QuantidadeTotalCarro int `json:"quantidade_total_carro"`
}