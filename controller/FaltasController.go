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

type FaltaController struct {
	service *service.FaltaService
}

func NewFaltaController(service *service.FaltaService) *FaltaController {
	return &FaltaController{
		service: service,
	}
}

// Save cadastra uma nova falta.
func (c *FaltaController) Save(w http.ResponseWriter, r *http.Request) {
	var falta entities.Falta

	if err := json.NewDecoder(r.Body).Decode(&falta); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err := c.service.Save(&falta)

	if errors.Is(err, service.ErrProdutoObrigatorio) ||
		errors.Is(err, service.ErrVendedorObrigatorio) ||
		errors.Is(err, service.ErrQuantidadeFaltaInvalida) ||
		errors.Is(err, service.ErrQuantidadeTotalInvalida) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrProdutoNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(
			w,
			"Erro interno ao salvar falta",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(falta)
}

// FindAll retorna todas as faltas.
func (c *FaltaController) FindAll(w http.ResponseWriter, r *http.Request) {
	faltas, err := c.service.FindAll()

	if err != nil {
		http.Error(
			w,
			"Erro interno ao buscar faltas",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(faltas)
}

// FindByID busca uma falta pelo ID.
func (c *FaltaController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := getFaltaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	falta, err := c.service.FindByID(id)

	if errors.Is(err, service.ErrFaltaNaoEncontrada) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(
			w,
			"Erro interno ao buscar falta",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(falta)
}

// FindByVendedorID retorna as faltas de um vendedor.
func (c *FaltaController) FindByVendedorID(w http.ResponseWriter, r *http.Request) {
	vendedorIDString := r.URL.Query().Get("vendedor_id")

	vendedorID, err := strconv.Atoi(vendedorIDString)

	if err != nil || vendedorID <= 0 {
		http.Error(w, "ID do vendedor inválido", http.StatusBadRequest)
		return
	}

	faltas, err := c.service.FindByVendedorID(vendedorID)

	if err != nil {
		http.Error(
			w,
			"Erro interno ao buscar faltas",
			http.StatusInternalServerError,
		)
		return
	}

	json.NewEncoder(w).Encode(faltas)
}

// Update atualiza uma falta.
func (c *FaltaController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getFaltaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var falta entities.Falta

	if err := json.NewDecoder(r.Body).Decode(&falta); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, falta)

	if errors.Is(err, service.ErrProdutoObrigatorio) ||
		errors.Is(err, service.ErrVendedorObrigatorio) ||
		errors.Is(err, service.ErrQuantidadeFaltaInvalida) ||
		errors.Is(err, service.ErrQuantidadeTotalInvalida) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrProdutoNaoEncontrado) ||
		errors.Is(err, service.ErrFaltaNaoEncontrada) {

		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(
			w,
			"Erro interno ao atualizar falta",
			http.StatusInternalServerError,
		)
		return
	}

	falta.ID = id

	json.NewEncoder(w).Encode(falta)
}

// Delete exclui uma falta.
func (c *FaltaController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getFaltaID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)

	if errors.Is(err, service.ErrFaltaNaoEncontrada) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(
			w,
			"Erro interno ao excluir falta",
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getFaltaID extrai o ID da URL.
func getFaltaID(r *http.Request) (int, error) {
	idString := strings.TrimPrefix(
		r.URL.Path,
		"/faltas/",
	)

	return strconv.Atoi(idString)
}