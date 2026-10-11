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

type RotaController struct {
	service *service.RotaService
}

func NewRotaController(service *service.RotaService) *RotaController {
	return &RotaController{
		service: service,
	}
}

// Save cadastra uma nova rota.
func (c *RotaController) Save(w http.ResponseWriter, r *http.Request) {
	var rota entities.Rota

	if err := json.NewDecoder(r.Body).Decode(&rota); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err := c.service.Save(&rota)

	if errors.Is(err, service.ErrCidadeObrigatoria) ||
		errors.Is(err, service.ErrBairroObrigatorio) ||
		errors.Is(err, service.ErrVendedorObrigatorio) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao salvar rota", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(rota)
}

// FindAll retorna todas as rotas.
func (c *RotaController) FindAll(w http.ResponseWriter, r *http.Request) {
	rotas, err := c.service.FindAll()

	if err != nil {
		http.Error(w, "Erro interno ao buscar rotas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rotas)
}

// FindByID busca uma rota pelo ID.
func (c *RotaController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := getRotaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	rota, err := c.service.FindByID(id)

	if errors.Is(err, service.ErrRotaNaoEncontrada) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao buscar rota", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rota)
}

// FindByVendedorID retorna as rotas de um vendedor.
func (c *RotaController) FindByVendedorID(w http.ResponseWriter, r *http.Request) {
	vendedorIDString := r.URL.Query().Get("vendedor_id")

	vendedorID, err := strconv.Atoi(vendedorIDString)

	if err != nil || vendedorID <= 0 {
		http.Error(w, "ID do vendedor inválido", http.StatusBadRequest)
		return
	}

	rotas, err := c.service.FindByVendedorID(vendedorID)

	if err != nil {
		http.Error(w, "Erro interno ao buscar rotas", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rotas)
}

// Update atualiza uma rota.
func (c *RotaController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getRotaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var rota entities.Rota

	if err := json.NewDecoder(r.Body).Decode(&rota); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, rota)

	if errors.Is(err, service.ErrCidadeObrigatoria) ||
		errors.Is(err, service.ErrBairroObrigatorio) ||
		errors.Is(err, service.ErrVendedorObrigatorio) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrRotaNaoEncontrada) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao atualizar rota", http.StatusInternalServerError)
		return
	}

	rota.ID = id
	json.NewEncoder(w).Encode(rota)
}

// Delete exclui uma rota.
func (c *RotaController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getRotaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)

	if errors.Is(err, service.ErrRotaNaoEncontrada) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao excluir rota", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getRotaID extrai o ID da URL.
func getRotaID(r *http.Request) (int, error) {
	idString := strings.TrimPrefix(r.URL.Path, "/rotas/")
	return strconv.Atoi(idString)
}