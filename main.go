package main

import (
	"fmt"
	"net/http"

	"crud-go/controller"
	"crud-go/database"
	"crud-go/repository"
	"crud-go/service"
)

// Middleware responsável por liberar o acesso do frontend à API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {

	// =========================
	// BANCO DE DADOS
	// =========================

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	// =========================
	// PRODUTO
	// =========================

	produtoRepository := repository.NewProdutoRepository(db)
	produtoService := service.NewProdutoService(produtoRepository)
	produtoController := controller.NewProdutoController(produtoService)

	// =========================
	// VENDEDOR
	// =========================

	vendedorRepository := repository.NewVendedorRepository(db)
	vendedorService := service.NewVendedorService(vendedorRepository)
	vendedorController := controller.NewVendedorController(vendedorService)

	// =========================
	// ROTA
	// =========================

	rotaRepository := repository.NewRotaRepository(db)
	rotaService := service.NewRotaService(rotaRepository)
	rotaController := controller.NewRotaController(rotaService)

	// =========================
	// FALTA
	// =========================

	faltaRepository := repository.NewFaltaRepository(db)
	faltaService := service.NewFaltaService(
	faltaRepository,
	produtoRepository,
)
	faltaController := controller.NewFaltaController(faltaService)

	// =========================
	// ROTAS DE PRODUTO
	// =========================

	http.HandleFunc("/produtos", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			produtoController.Save(w, r)

		case http.MethodGet:
			produtoController.FindAll(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	http.HandleFunc("/produtos/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			produtoController.FindByID(w, r)

		case http.MethodPut:
			produtoController.Update(w, r)

		case http.MethodDelete:
			produtoController.Delete(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	// =========================
	// ROTAS DE VENDEDOR
	// =========================

	http.HandleFunc("/vendedores", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			vendedorController.Save(w, r)

		case http.MethodGet:
			vendedorController.FindAll(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	http.HandleFunc("/vendedores/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			vendedorController.FindByID(w, r)

		case http.MethodPut:
			vendedorController.Update(w, r)

		case http.MethodDelete:
			vendedorController.Delete(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	// =========================
	// ROTAS DE ROTA
	// =========================

	http.HandleFunc("/rotas", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			rotaController.Save(w, r)

		case http.MethodGet:

			// Se houver vendedor_id, retorna apenas
			// as rotas daquele vendedor.
			if r.URL.Query().Get("vendedor_id") != "" {
				rotaController.FindByVendedorID(w, r)
				return
			}

			rotaController.FindAll(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	http.HandleFunc("/rotas/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			rotaController.FindByID(w, r)

		case http.MethodPut:
			rotaController.Update(w, r)

		case http.MethodDelete:
			rotaController.Delete(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	// =========================
	// ROTAS DE FALTA
	// =========================

	http.HandleFunc("/faltas", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodPost:
			faltaController.Save(w, r)

		case http.MethodGet:

			// Se houver vendedor_id, retorna apenas
			// as faltas daquele vendedor.
			if r.URL.Query().Get("vendedor_id") != "" {
				faltaController.FindByVendedorID(w, r)
				return
			}

			faltaController.FindAll(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	http.HandleFunc("/faltas/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {

		case http.MethodGet:
			faltaController.FindByID(w, r)

		case http.MethodPut:
			faltaController.Update(w, r)

		case http.MethodDelete:
			faltaController.Delete(w, r)

		default:
			http.Error(
				w,
				"Método não permitido",
				http.StatusMethodNotAllowed,
			)
		}
	})

	// =========================
	// SERVIDOR
	// =========================

	fmt.Println("Servidor rodando em http://localhost:8080")

	err = http.ListenAndServe(
		":8080",
		corsMiddleware(http.DefaultServeMux),
	)

	if err != nil {
		panic(err)
	}
}