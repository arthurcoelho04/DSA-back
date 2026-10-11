package entities

type Vendedor struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Nome string `json:"nome"`
}