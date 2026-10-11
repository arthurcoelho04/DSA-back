package service

import (
	"crud-go/entities"
	"errors"
	"testing"
)

// Testa se um produto sem nome é rejeitado.
func TestValidarProdutoNomeVazio(t *testing.T) {
	produto := entities.Produto{
		Nome: "",
	}

	err := validarProduto(produto)

	if !errors.Is(err, ErrNomeObrigatorio) {
		t.Errorf("esperado ErrNomeObrigatorio, recebido: %v", err)
	}
}

// Testa se um produto válido passa pela validação sem erro.
func TestValidarProdutoValido(t *testing.T) {
	produto := entities.Produto{
		Nome: "Teclado",
	}

	err := validarProduto(produto)

	if err != nil {
		t.Errorf("não era esperado erro, recebido: %v", err)
	}
}