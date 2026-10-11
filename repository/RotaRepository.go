package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type RotaRepository struct {
	db *gorm.DB
}

func NewRotaRepository(db *gorm.DB) *RotaRepository {
	return &RotaRepository{
		db: db,
	}
}

// Salva uma nova rota.
func (r *RotaRepository) Save(rota *entities.Rota) error {
	return r.db.Create(rota).Error
}

// Retorna todas as rotas cadastradas.
func (r *RotaRepository) FindAll() ([]entities.Rota, error) {
	var rotas []entities.Rota

	err := r.db.
		Preload("Vendedor").
		Find(&rotas).Error

	return rotas, err
}

// Busca uma rota pelo ID.
func (r *RotaRepository) FindByID(id int) (entities.Rota, error) {
	var rota entities.Rota

	err := r.db.
		Preload("Vendedor").
		First(&rota, id).Error

	return rota, err
}

// Busca todas as rotas de um vendedor.
func (r *RotaRepository) FindByVendedorID(vendedorID int) ([]entities.Rota, error) {
	var rotas []entities.Rota

	err := r.db.
		Preload("Vendedor").
		Where("vendedor_id = ?", vendedorID).
		Find(&rotas).Error

	return rotas, err
}

// Atualiza uma rota.
func (r *RotaRepository) Update(id int, rota entities.Rota) error {
	result := r.db.Model(&entities.Rota{}).
		Where("id = ?", id).
		Select(
			"Cidade",
			"Bairro",
			"VendedorID",
		).
		Updates(rota)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Exclui uma rota.
func (r *RotaRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Rota{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}