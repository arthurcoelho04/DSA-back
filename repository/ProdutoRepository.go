package repository

import (
	"crud-go/entities"

	"gorm.io/gorm"
)

type ProdutoRepository struct {
	db *gorm.DB										// possui uma dependencia do tipo grom.DB
}

func NewProdutoRepository(db *gorm.DB) *ProdutoRepository {
	return &ProdutoRepository{										// "construtor"
		db: db,
	}
}

func (r *ProdutoRepository) Save(produto *entities.Produto) error {  //Ponteiro
	return r.db.Create(produto).Error
}

func (r *ProdutoRepository) FindAll() ([]entities.Produto, error) {
	var produtos []entities.Produto

	err := r.db.Find(&produtos).Error  							//endereço de momoria do produto

	return produtos, err
}

func (r *ProdutoRepository) FindByID(id int) (entities.Produto, error) {
	var produto entities.Produto

	err := r.db.First(&produto, id).Error 					//Grom

	return produto, err
}

func (r *ProdutoRepository) Delete(id int) error {
	result := r.db.Delete(&entities.Produto{}, id)			//Grom

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

// Atualiza um produto pelo ID, incluindo campos com valor zero.
func (r *ProdutoRepository) Update(id int, produto entities.Produto) error {
	result := r.db.Model(&entities.Produto{}).
		Where("id = ?", id). 						//atualize somente onde o ID for igual ao id recebido.
		Select("Nome", "QuantidadeTotalCarro").                    //Quero atualizar especificamente os campos Nome
		Updates(produto)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}