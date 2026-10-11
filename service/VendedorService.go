package service

import (
	"crud-go/entities"
	"crud-go/repository"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type VendedorService struct {
	repository *repository.VendedorRepository
}

func NewVendedorService(repository *repository.VendedorRepository) *VendedorService {
	return &VendedorService{
		repository: repository,
	}
}

func validarVendedor(vendedor entities.Vendedor) error {
	if strings.TrimSpace(vendedor.Nome) == "" {
		return ErrNomeVendedorObrigatorio
	}

	return nil
}

func (s *VendedorService) Save(vendedor *entities.Vendedor) error {
	if err := validarVendedor(*vendedor); err != nil {
		return err
	}

	return s.repository.Save(vendedor)
}

func (s *VendedorService) FindAll() ([]entities.Vendedor, error) {
	return s.repository.FindAll()
}

func (s *VendedorService) FindByID(id int) (entities.Vendedor, error) {
	vendedor, err := s.repository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Vendedor{}, ErrVendedorNaoEncontrado
	}

	if err != nil {
		return entities.Vendedor{}, err
	}

	return vendedor, nil
}

func (s *VendedorService) Update(id int, vendedorAtualizado entities.Vendedor) error {
	if err := validarVendedor(vendedorAtualizado); err != nil {
		return err
	}

	err := s.repository.Update(id, vendedorAtualizado)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrVendedorNaoEncontrado
	}

	return err
}

func (s *VendedorService) Delete(id int) error {
	err := s.repository.Delete(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrVendedorNaoEncontrado
	}

	return err
}