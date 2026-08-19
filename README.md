# GoJob Opportunities API

<p align="center">
  <img src="./assets/GopportunitiesHeader.svg" alt="GoJob Header">
</p>

API REST desenvolvida em **Go (Golang)** para gerenciamento de oportunidades de emprego.

O projeto foi desenvolvido com foco no aprendizado de desenvolvimento backend em Go, trabalhando conceitos como criação de APIs REST, organização de projetos, persistência de dados, documentação de endpoints, testes automatizados e containerização.

## 🚀 Funcionalidades

* Cadastro de oportunidades de emprego
* Consulta de oportunidades
* Busca de oportunidades por diferentes critérios
* Atualização de oportunidades
* Exclusão de oportunidades
* Validação dos dados recebidos pela API
* Persistência dos dados em banco SQLite
* Documentação interativa com Swagger
* Testes automatizados
* Execução utilizando Docker e Docker Compose

## 🛠️ Tecnologias utilizadas

* **Go (Golang)** — desenvolvimento do backend
* **Gin** — criação e gerenciamento das rotas HTTP
* **GORM** — comunicação com o banco de dados
* **SQLite** — armazenamento dos dados
* **Swagger** — documentação e teste dos endpoints
* **Docker** — containerização da aplicação
* **Docker Compose** — gerenciamento dos containers

## 📁 Estrutura do projeto

O projeto utiliza uma organização baseada em pacotes, buscando separar as responsabilidades da aplicação e facilitar sua manutenção e evolução.

```text
.
├── assets/
├── controllers/
├── database/
├── models/
├── tests/
├── docs/
├── handler/
├── router/
├── schemas/
├── docker-compose.yml
├── Dockerfile
├── Makefile
├── go.mod
└── main.go
```

> A estrutura pode variar conforme a evolução do projeto.

## ⚙️ Como executar

### Pré-requisitos

Antes de iniciar, certifique-se de ter instalado:

* Go
* Git
* Docker (opcional)

### Executando localmente

Clone o repositório:

```bash
git clone <URL_DO_SEU_REPOSITORIO>
cd <NOME_DO_REPOSITORIO>
```

Baixe as dependências:

```bash
go mod download
```

Execute a aplicação:

```bash
go run .
```

Por padrão, a API será disponibilizada na porta `8080`.

## 📚 Documentação da API

Após iniciar a aplicação, a documentação do Swagger pode ser acessada em:

```text
http://localhost:8080/swagger/index.html
```

Através da interface do Swagger é possível visualizar os endpoints disponíveis e realizar requisições diretamente pelo navegador.

## 🔨 Comandos do Makefile

O projeto possui um `Makefile` para facilitar tarefas comuns durante o desenvolvimento.

| Comando              | Descrição                                |
| -------------------- | ---------------------------------------- |
| `make run`           | Executa a aplicação                      |
| `make run-with-docs` | Gera a documentação e inicia a aplicação |
| `make build`         | Compila a aplicação                      |
| `make test`          | Executa os testes                        |
| `make docs`          | Gera a documentação do Swagger           |
| `make clean`         | Remove arquivos gerados                  |

Exemplo:

```bash
make run
```

## 🐳 Docker

Também é possível executar a aplicação utilizando Docker.

Para criar a imagem:

```bash
docker build -t gojob-api .
```

Para executar o container:

```bash
docker run -p 8080:8080 gojob-api
```

Caso utilize Docker Compose:

```bash
docker compose build
docker compose up
```

Para interromper os serviços:

```bash
docker compose down
```

## 🔌 Principais endpoints

| Método   | Endpoint             | Descrição                 |
| -------- | -------------------- | ------------------------- |
| `GET`    | `/opportunities`     | Lista oportunidades       |
| `GET`    | `/opportunities/:id` | Busca uma oportunidade    |
| `POST`   | `/opportunities`     | Cria uma oportunidade     |
| `PUT`    | `/opportunities/:id` | Atualiza uma oportunidade |
| `DELETE` | `/opportunities/:id` | Remove uma oportunidade   |

> Os endpoints podem variar conforme a implementação atual da API.

## 🧪 Testes

Os testes podem ser executados utilizando:

```bash
go test ./...
```

O objetivo dos testes é verificar o comportamento dos principais componentes da aplicação e reduzir possíveis regressões durante o desenvolvimento.

## 🎯 Objetivo do projeto

Este projeto faz parte dos meus estudos em **desenvolvimento backend com Go**.

Durante o desenvolvimento, foram praticados conceitos importantes como:

* Desenvolvimento de APIs REST
* Routing
* HTTP methods e status codes
* CRUD
* ORM
* Persistência de dados
* Estruturação de projetos Go
* Testes automatizados
* Documentação de APIs
* Docker
* Boas práticas de desenvolvimento backend

## 📌 Próximos passos

Algumas melhorias que podem ser implementadas futuramente:

* [ ] Adicionar autenticação e autorização
* [ ] Implementar paginação
* [ ] Adicionar filtros mais avançados
* [ ] Melhorar a cobertura de testes
* [ ] Utilizar PostgreSQL em ambiente de produção
* [ ] Implementar CI/CD
* [ ] Adicionar observabilidade e logs estruturados

## 📄 Licença

Este projeto está disponível sob a licença definida no arquivo `LICENSE`.

---

Desenvolvido por **Guilherme Freires** como projeto de estudo em desenvolvimento backend com Go.
