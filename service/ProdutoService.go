package service

import (
	"crud-go/entities"
	"crud-go/repository"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type ProdutoService struct {
	repository *repository.ProdutoRepository
}

func NewProdutoService(repository *repository.ProdutoRepository) *ProdutoService {
	return &ProdutoService{
		repository: repository,
	}
}

func validarProduto(produto entities.Produto) error {

	if strings.TrimSpace(produto.Nome) == "" {
		return ErrNomeObrigatorio
	}

	if produto.QuantidadeTotalCarro < 0 {
		return ErrQuantidadeTotalCarroInvalida
	}

	return nil
}

func (s *ProdutoService) Save(produto *entities.Produto) error {

	if err := validarProduto(*produto); err != nil {
		return err
	}

	return s.repository.Save(produto)
}

func (s *ProdutoService) FindAll() ([]entities.Produto, error) {
	return s.repository.FindAll()
}

func (s *ProdutoService) FindByID(id int) (entities.Produto, error) {

	produto, err := s.repository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Produto{}, ErrProdutoNaoEncontrado
	}

	if err != nil {
		return entities.Produto{}, err
	}

	return produto, nil
}

func (s *ProdutoService) Delete(id int) error {

	err := s.repository.Delete(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrProdutoNaoEncontrado
	}

	return err
}

func (s *ProdutoService) Update(
	id int,
	produtoAtualizado entities.Produto,
) error {

	if err := validarProduto(produtoAtualizado); err != nil {
		return err
	}

	err := s.repository.Update(id, produtoAtualizado)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrProdutoNaoEncontrado
	}

	return err
}