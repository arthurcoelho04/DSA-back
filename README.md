# DSA — Backend

API REST desenvolvida em **Go (Golang)**, responsável pelo gerenciamento de dados e pelas regras de negócio do sistema DSA.

## Funcionalidades

- **Gerenciamento de Produtos:** Cadastro, consulta, edição e exclusão de produtos.
- **Gerenciamento de Vendedores:** Controle e administração dos vendedores.
- **Controle de Faltas:** Registro e acompanhamento de faltas de produtos.
- **API REST:** Comunicação entre o frontend e o backend.
- **Persistência de Dados:** Armazenamento e gerenciamento das informações no banco de dados.

## Tecnologias Utilizadas

- **Go (Golang):** Desenvolvimento do backend.
- **GORM:** Mapeamento objeto-relacional (ORM).
- **PostgreSQL:** Banco de dados relacional.
- **API REST:** Integração entre aplicações.

## Estrutura do Projeto

- `controller/` — Gerenciamento das requisições HTTP.
- `service/` — Regras de negócio.
- `repository/` — Acesso e manipulação dos dados.
- `entities/` — Modelos das entidades.
- `database/` — Configuração e conexão com o banco.
- `main.go` — Inicialização da aplicação.

## Como Executar

1. Clone o repositório.
2. Configure a conexão com o PostgreSQL.
3. Instale as dependências com `go mod download`.
4. Execute a aplicação com `go run .`.

## Frontend

A interface web está disponível no repositório:

[DSA — Frontend](https://github.com/arthurcoelho04/DSA-FRONT)
