package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type FaltaRepository struct {
	db *gorm.DB
}

func NewFaltaRepository(db *gorm.DB) *FaltaRepository {
	return &FaltaRepository{
		db: db,
	}
}

// Salva uma nova falta.
func (r *FaltaRepository) Save(falta *entities.Falta) error {
	return r.db.Create(falta).Error
}

// Retorna todas as faltas cadastradas,
// carregando também Produto e Vendedor.
func (r *FaltaRepository) FindAll() ([]entities.Falta, error) {
	var faltas []entities.Falta

	err := r.db.
		Preload("Produto").
		Preload("Vendedor").
		Find(&faltas).Error

	return faltas, err
}

// Busca uma falta pelo ID.
func (r *FaltaRepository) FindByID(id int) (entities.Falta, error) {
	var falta entities.Falta

	err := r.db.
		Preload("Produto").
		Preload("Vendedor").
		First(&falta, id).Error

	return falta, err
}

// Busca todas as faltas de um vendedor.
func (r *FaltaRepository) FindByVendedorID(vendedorID int) ([]entities.Falta, error) {
	var faltas []entities.Falta

	err := r.db.
		Preload("Produto").
		Preload("Vendedor").
		Where("vendedor_id = ?", vendedorID).
		Find(&faltas).Error

	return faltas, err
}

// Atualiza uma falta.
func (r *FaltaRepository) Update(id int, falta entities.Falta) error {
	result := r.db.Model(&entities.Falta{}).
		Where("id = ?", id).
		Select(
			"QuantidadeFalta",
			"QuantidadeTotal",
			"Data",
			"ProdutoID",
			"VendedorID",
		).
		Updates(falta)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Exclui uma falta.
func (r *FaltaRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Falta{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
