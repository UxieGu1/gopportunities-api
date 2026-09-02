<p align="center">
  <img src="./assets/GopportunitiesHeader.svg" alt="GoJob Header">
</p>

# GoJob Opportunities API

API REST completa desenvolvida em Go (Golang) para gerenciamento de um ecossistema de vagas de emprego, conectando empresas, oportunidades, usuários e candidatos.

O projeto evoluiu de um simples CRUD para uma arquitetura robusta e em camadas, focando no aprendizado avançado de desenvolvimento backend. A aplicação implementa injeção de dependência, padrão Repository, separação clara de regras de negócio (Services) e roteamento limpo utilizando o framework Gin.

## 📋 Índice

- [Funcionalidades](#-funcionalidades)
- [Tecnologias Utilizadas](#️-tecnologias-utilizadas)
- [Estrutura do Projeto](#-estrutura-do-projeto)
- [Como Executar](#-como-executar)
- [Documentação da API](#-documentação-da-api-swagger)
- [Principais Endpoints](#-principais-endpoints-api-v1)
- [Makefile](#-makefile)
- [Docker](#-docker)
- [Próximos Passos](#-próximos-passos)
- [Licença](#-licença)

## 🚀 Funcionalidades

O sistema foi expandido e agora gerencia cinco domínios principais:

- **Openings (Vagas)**: Cadastro, listagem e detalhamento de oportunidades de emprego
- **Companies (Empresas)**: Gerenciamento das empresas que publicam as vagas
- **Users (Usuários)**: Base de autenticação e perfis da plataforma
- **Candidates (Candidatos)**: Perfis profissionais vinculados aos usuários, contendo currículos e habilidades
- **Applications (Candidaturas)**: Fluxo de aplicação de candidatos às vagas, com rastreamento de status (APPLIED, REVIEWING, etc)

## 🛠️ Tecnologias Utilizadas

- **Go (Golang)** — Desenvolvimento do backend
- **Gin** — Roteamento e middlewares HTTP
- **GORM** — ORM para mapeamento das entidades e relações
- **SQLite** — Banco de dados relacional embarcado (desenvolvimento)
- **Swagger** — Documentação interativa da API
- **Docker & Compose** — Containerização e orquestração de ambiente

## 📁 Estrutura do Projeto

O projeto adota uma arquitetura em camadas (Clean Architecture inspired), separando responsabilidades por domínio e facilitando a manutenção, testes e evolução da base de código.

````
.
├── cmd/
│   └── api/
│       └── main.go               # Ponto de entrada da aplicação
├── internal/
│   ├── config/                   # Configurações globais (Logger, DB)
│   ├── handler/                  # Controladores HTTP (Requests, Responses)
│   │   ├── application/
│   │   ├── candidate/
│   │   ├── company/
│   │   ├── opening/
│   │   └── user/
│   ├── repository/               # Camada de acesso a dados (GORM/SQLite)
│   ├── router/                   # Definição e agrupamento de rotas (Gin)
│   ├── schemas/                  # Entidades do banco e DTOs
│   └── service/                  # Regras de negócio e casos de uso
├── docs/                         # Documentação gerada pelo Swagger
├── .env.example
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── go.mod
w```

## ⚙️ Como Executar

### Pré-requisitos

- Go 1.22+
- Git
- Docker (opcional)

### Executando Localmente

1. Clone o repositório:

```bash
git clone github.com/UxieGu1/gopportunities-api
cd gopportunities-api
````

2. Baixe as dependências:

```bash
go mod download
```

3. Execute a aplicação:

```bash
go run ./cmd/api
```

A API estará disponível em `http://localhost:8081`.

## 📚 Documentação da API (Swagger)

A documentação interativa pode ser acessada com o servidor rodando em:

```
http://localhost:8081/swagger/index.html
```

## 🔌 Principais Endpoints (API v1)

A API é prefixada por `/api/v1`. Abaixo estão os recursos disponíveis:

| Recurso          | Criação (POST)  | Leitura (GET)                         | Atualização (PUT)   | Exclusão (DELETE)   |
| ---------------- | --------------- | ------------------------------------- | ------------------- | ------------------- | -------------- |
| **Openings**     | `/openings`     | `/openings` e `/openings/:id`         | `/openings/:id`     | `/openings/:id`     |
| **Companies**    | `/companies`    | `/companies` e `/companies/:id`       | `/companies/:id`    | `/companies/:id`    |
| **Users**        | `/users`        | `/users` e `/users/:id`               | `/users/:id`        | `/users/:id`        |
| **Candidates**   | `/candidates`   | `/candidates` e `/candidates/:id`     | `/candidates/:id`   | `/candidates/:id`   |
| **Applications** | `/applications` | `/applications` e `/applications/:id` | `/applications/:id` | `/applications/:id` | ## 🔨 Makefile |

Utilitário de linha de comando para acelerar o desenvolvimento:

| Comando              | Descrição                                  |
| -------------------- | ------------------------------------------ | ------------ |
| `make run`           | Executa a aplicação                        |
| `make run-with-docs` | Gera a documentação Swagger e inicia a API |
| `make build`         | Compila o binário da aplicação             |
| `make test`          | Executa os testes unitários                |
| `make docs`          | Atualiza a documentação do Swagger         |
| `make clean`         | Limpa o binário gerado                     | ## 🐳 Docker |

Para executar a aplicação em um ambiente isolado:

```bash
# Build e execução via Docker Compose
docker compose build
docker compose up -d

# Para encerrar a aplicação
docker compose down
```

## 📌 Próximos Passos

Evoluções planejadas para a maturidade do projeto:

- [ ] Implementar criptografia de senhas para Users (bcrypt)
- [ ] Adicionar autenticação e autorização via JWT
- [ ] Escrever testes unitários para a camada de Service usando Mocks
- [ ] Migrar o banco de dados de SQLite para PostgreSQL no ambiente de produção
- [ ] Implementar paginação e filtros nas rotas de listagem (GET)
- [ ] Adicionar pipeline de CI/CD (GitHub Actions)## 📄 Licença

Este projeto está disponível sob a licença definida no arquivo [LICENSE](LICENSE).

---

Desenvolvido por **Guilherme Freires** como projeto prático para domínio da linguagem Go e arquitetura de software.
