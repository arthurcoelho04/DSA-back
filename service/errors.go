package service

import "errors"

var (
	ErrNomeObrigatorio      = errors.New("nome do produto é obrigatório")
	ErrProdutoNaoEncontrado = errors.New("produto não encontrado")

	ErrNomeVendedorObrigatorio  = errors.New("nome do vendedor obrigatório")
	ErrVendedorNaoEncontrado    = errors.New("vendedor não encontrado")

	ErrProdutoObrigatorio      = errors.New("produto obrigatório")
	ErrVendedorObrigatorio     = errors.New("vendedor obrigatório")
	ErrQuantidadeFaltaInvalida = errors.New("quantidade faltante inválida")
	ErrQuantidadeTotalInvalida = errors.New("quantidade total inválida")
	ErrFaltaNaoEncontrada      = errors.New("falta não encontrada")
	ErrQuantidadeTotalCarroInvalida = errors.New("quantidade total no carro inválida")

	ErrCidadeObrigatoria    = errors.New("cidade obrigatória")
	ErrBairroObrigatorio    = errors.New("bairro obrigatório")
	ErrRotaNaoEncontrada    = errors.New("rota não encontrada")
)