package entities

type Rota struct {
	ID int `json:"id" gorm:"primaryKey"`

	Cidade string `json:"cidade"`
	Bairro string `json:"bairro"`

	VendedorID int      `json:"vendedor_id"`
	Vendedor   Vendedor `json:"vendedor" gorm:"foreignKey:VendedorID"`
}