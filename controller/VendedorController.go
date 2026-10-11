package controller

import (
	"crud-go/entities"
	"crud-go/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type VendedorController struct {
	service *service.VendedorService
}

func NewVendedorController(service *service.VendedorService) *VendedorController {
	return &VendedorController{
		service: service,
	}
}

// Save cadastra um novo vendedor.
func (c *VendedorController) Save(w http.ResponseWriter, r *http.Request) {
	var vendedor entities.Vendedor

	if err := json.NewDecoder(r.Body).Decode(&vendedor); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := c.service.Save(&vendedor); err != nil {
		if errors.Is(err, service.ErrNomeVendedorObrigatorio) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Erro interno ao salvar vendedor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vendedor)
}

// FindAll retorna todos os vendedores cadastrados.
func (c *VendedorController) FindAll(w http.ResponseWriter, r *http.Request) {
	vendedores, err := c.service.FindAll()

	if err != nil {
		http.Error(w, "Erro interno ao buscar vendedores", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vendedores)
}

// FindByID busca um vendedor pelo ID.
func (c *VendedorController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := getVendedorID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	vendedor, err := c.service.FindByID(id)

	if errors.Is(err, service.ErrVendedorNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao buscar vendedor", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(vendedor)
}

// Update atualiza os dados de um vendedor.
func (c *VendedorController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getVendedorID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var vendedor entities.Vendedor

	if err := json.NewDecoder(r.Body).Decode(&vendedor); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, vendedor)

	if errors.Is(err, service.ErrNomeVendedorObrigatorio) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrVendedorNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao atualizar vendedor", http.StatusInternalServerError)
		return
	}

	vendedor.ID = id
	json.NewEncoder(w).Encode(vendedor)
}

// Delete exclui um vendedor pelo ID.
func (c *VendedorController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getVendedorID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)

	if errors.Is(err, service.ErrVendedorNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao excluir vendedor", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getVendedorID extrai o ID da URL.
func getVendedorID(r *http.Request) (int, error) {
	idString := strings.TrimPrefix(r.URL.Path, "/vendedores/")
	return strconv.Atoi(idString)
}