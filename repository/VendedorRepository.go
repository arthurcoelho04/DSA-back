package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type VendedorRepository struct {
	db *gorm.DB
}

func NewVendedorRepository(db *gorm.DB) *VendedorRepository {
	return &VendedorRepository{
		db: db,
	}
}

func (r *VendedorRepository) Save(vendedor *entities.Vendedor) error {
	return r.db.Create(vendedor).Error
}

func (r *VendedorRepository) FindAll() ([]entities.Vendedor, error) {
	var vendedores []entities.Vendedor

	err := r.db.Find(&vendedores).Error

	return vendedores, err
}

func (r *VendedorRepository) FindByID(id int) (entities.Vendedor, error) {
	var vendedor entities.Vendedor

	err := r.db.First(&vendedor, id).Error

	return vendedor, err
}

func (r *VendedorRepository) Update(id int, vendedor entities.Vendedor) error {
	result := r.db.Model(&entities.Vendedor{}).
		Where("id = ?", id).
		Select("Nome").
		Updates(vendedor)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}


func (r *VendedorRepository) Delete(id int) error {
    return r.db.Transaction(func(tx *gorm.DB) error {

        // 1. Excluir as rotas vinculadas ao vendedor
        if err := tx.Where("vendedor_id = ?", id).
            Delete(&entities.Rota{}).Error; err != nil {
            return err
        }

        // 2. Excluir o vendedor
        result := tx.Delete(&entities.Vendedor{}, id)

        if result.Error != nil {
            return result.Error
        }

        // 3. Verificar se o vendedor existe
        if result.RowsAffected == 0 {
            return gorm.ErrRecordNotFound
        }

        return nil
    })
}
