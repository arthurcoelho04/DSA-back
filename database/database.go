package database

import (
	"crud-go/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=1234 dbname=dsa_faltas port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
	&entities.Produto{},
	&entities.Vendedor{},
	&entities.Rota{},
	&entities.Falta{},
)

	if err != nil {
		return nil, err
	}

	fmt.Println("Conectado ao banco de dados!")

	return db, nil
}