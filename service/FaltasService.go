package service

import (
	"crud-go/entities"
	"crud-go/repository"
	"errors"

	"gorm.io/gorm"
)

type FaltaService struct {
	repository        *repository.FaltaRepository
	produtoRepository *repository.ProdutoRepository
}

func NewFaltaService(
	repository *repository.FaltaRepository,
	produtoRepository *repository.ProdutoRepository,
) *FaltaService {

	return &FaltaService{
		repository:        repository,
		produtoRepository: produtoRepository,
	}
}

// Valida somente os dados que o usuário informa.
func validarFalta(falta entities.Falta) error {

	if falta.ProdutoID <= 0 {
		return ErrProdutoObrigatorio
	}

	if falta.VendedorID <= 0 {
		return ErrVendedorObrigatorio
	}

	if falta.QuantidadeFalta <= 0 {
		return ErrQuantidadeFaltaInvalida
	}

	return nil
}

// Busca o produto e coloca automaticamente
// a quantidade total do carro na falta.
func (s *FaltaService) preencherQuantidadeTotal(
	falta *entities.Falta,
) error {

	produto, err := s.produtoRepository.FindByID(falta.ProdutoID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrProdutoNaoEncontrado
	}

	if err != nil {
		return err
	}

	if produto.QuantidadeTotalCarro <= 0 {
		return ErrQuantidadeTotalInvalida
	}

	falta.QuantidadeTotal = produto.QuantidadeTotalCarro

	return nil
}

func (s *FaltaService) Save(falta *entities.Falta) error {

	if err := validarFalta(*falta); err != nil {
		return err
	}

	// Quantidade total vem automaticamente do produto
	if err := s.preencherQuantidadeTotal(falta); err != nil {
		return err
	}

	return s.repository.Save(falta)
}

func (s *FaltaService) FindAll() ([]entities.Falta, error) {
	return s.repository.FindAll()
}

func (s *FaltaService) FindByID(id int) (entities.Falta, error) {

	falta, err := s.repository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Falta{}, ErrFaltaNaoEncontrada
	}

	if err != nil {
		return entities.Falta{}, err
	}

	return falta, nil
}

func (s *FaltaService) FindByVendedorID(
	vendedorID int,
) ([]entities.Falta, error) {

	if vendedorID <= 0 {
		return nil, ErrVendedorObrigatorio
	}

	return s.repository.FindByVendedorID(vendedorID)
}

func (s *FaltaService) Update(
	id int,
	faltaAtualizada entities.Falta,
) error {

	if err := validarFalta(faltaAtualizada); err != nil {
		return err
	}

	// Recalcula automaticamente conforme o produto escolhido
	if err := s.preencherQuantidadeTotal(&faltaAtualizada); err != nil {
		return err
	}

	err := s.repository.Update(id, faltaAtualizada)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrFaltaNaoEncontrada
	}

	return err
}

func (s *FaltaService) Delete(id int) error {

	err := s.repository.Delete(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrFaltaNaoEncontrada
	}

	return err
}