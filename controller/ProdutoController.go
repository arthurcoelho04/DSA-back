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

type ProdutoController struct {
	service *service.ProdutoService
}

func NewProdutoController(service *service.ProdutoService) *ProdutoController {
	return &ProdutoController{service: service}
}

// Save recebe os dados do produto e realiza o cadastro. (POST)
func (c *ProdutoController) Save(w http.ResponseWriter, r *http.Request) {
	var produto entities.Produto

	if err := json.NewDecoder(r.Body).Decode(&produto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	if err := c.service.Save(&produto); err != nil {
		if errors.Is(err, service.ErrNomeObrigatorio) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		http.Error(w, "Erro interno ao salvar produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(produto)
}

// FindAll retorna todos os produtos cadastrados. (GET)
func (c *ProdutoController) FindAll(w http.ResponseWriter, r *http.Request) {
	produtos, err := c.service.FindAll()

	if err != nil {
		http.Error(w, "Erro interno ao buscar produtos", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produtos)
}

// FindByID busca e retorna um produto através do seu ID.
func (c *ProdutoController) FindByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	produto, err := c.service.FindByID(id)

	if errors.Is(err, service.ErrProdutoNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao buscar produto", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(produto)
}

// Update recebe novos dados e atualiza um produto existente.
func (c *ProdutoController) Update(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var produto entities.Produto

	if err := json.NewDecoder(r.Body).Decode(&produto); err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.service.Update(id, produto)

		if errors.Is(err, service.ErrNomeObrigatorio) ||
		errors.Is(err, service.ErrQuantidadeTotalCarroInvalida) {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if errors.Is(err, service.ErrProdutoNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao atualizar produto", http.StatusInternalServerError)
		return
	}

	produto.ID = id
	json.NewEncoder(w).Encode(produto)
}

// Delete exclui um produto através do seu ID.
func (c *ProdutoController) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r)

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	err = c.service.Delete(id)

	if errors.Is(err, service.ErrProdutoNaoEncontrado) {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Erro interno ao excluir produto", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// getID extrai e converte o ID informado na URL.
func getID(r *http.Request) (int, error) {
	idString := strings.TrimPrefix(r.URL.Path, "/produtos/")
	return strconv.Atoi(idString)
}