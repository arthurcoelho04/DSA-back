package service

import (
	"crud-go/entities"
	"crud-go/repository"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type RotaService struct {
	repository *repository.RotaRepository
}

func NewRotaService(repository *repository.RotaRepository) *RotaService {
	return &RotaService{
		repository: repository,
	}
}

func validarRota(rota entities.Rota) error {

	if strings.TrimSpace(rota.Cidade) == "" {
		return ErrCidadeObrigatoria
	}

	if strings.TrimSpace(rota.Bairro) == "" {
		return ErrBairroObrigatorio
	}

	if rota.VendedorID <= 0 {
		return ErrVendedorObrigatorio
	}

	return nil
}

func (s *RotaService) Save(rota *entities.Rota) error {

	if err := validarRota(*rota); err != nil {
		return err
	}

	return s.repository.Save(rota)
}

func (s *RotaService) FindAll() ([]entities.Rota, error) {
	return s.repository.FindAll()
}

func (s *RotaService) FindByID(id int) (entities.Rota, error) {

	rota, err := s.repository.FindByID(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entities.Rota{}, ErrRotaNaoEncontrada
	}

	if err != nil {
		return entities.Rota{}, err
	}

	return rota, nil
}

func (s *RotaService) FindByVendedorID(vendedorID int) ([]entities.Rota, error) {

	if vendedorID <= 0 {
		return nil, ErrVendedorObrigatorio
	}

	return s.repository.FindByVendedorID(vendedorID)
}

func (s *RotaService) Update(id int, rotaAtualizada entities.Rota) error {

	if err := validarRota(rotaAtualizada); err != nil {
		return err
	}

	err := s.repository.Update(id, rotaAtualizada)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRotaNaoEncontrada
	}

	return err
}

func (s *RotaService) Delete(id int) error {

	err := s.repository.Delete(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrRotaNaoEncontrada
	}

	return err
}