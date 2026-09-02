.PHONY: default run run-with-docs build test docs clean help

# Nome do executável que será gerado
APP_NAME=gopportunities

# Tarefa padrão executada quando você digita apenas "make"
default: run

# Roda a aplicação em modo de desenvolvimento
run:
	@echo "🚀 Iniciando a API em modo dev..."
	@go run ./cmd/api

# Gera a documentação do Swagger (se você estiver usando swag init)
run-with-docs:
	@echo "📚 Atualizando documentação do Swagger..."
	@swag init -g cmd/api/main.go -o docs --parseDependency
	@go run ./cmd/api

# Gera apenas a documentação
docs:
	@echo "📚 Atualizando documentação do Swagger..."
	@swag init -g cmd/api/main.go -o docs --parseDependency

# Compila o projeto e gera o arquivo executável
build:
	@echo "🔨 Compilando o projeto..."
	@go build -o $(APP_NAME) ./cmd/api
	@echo "✅ Build concluído! Executável criado: $(APP_NAME)"

# Executa todos os testes do projeto
test:
	@echo "🧪 Rodando os testes..."
	@go test ./... -v


# Limpa os arquivos compilados (útil antes de subir para o GitHub ou refazer o build)
clean:
	@echo "🧹 Limpando o projeto..."
	@rm -f $(APP_NAME)
	@rm -rf ./docs
	@echo "✅ Limpeza concluída!"

# Comando de ajuda para listar as opções (basta digitar "make help")
help:
	@echo "Comandos disponíveis:"
	@echo "  make run    - Roda a aplicação (go run ./cmd/api)"
	@echo "  make build  - Compila o executável do projeto"
	@echo "  make test   - Executa todos os testes do repositório"
	@echo "  make docs   - Gera a documentação da API (swaggo)"
	@echo "  make clean  - Remove os binários compilados"