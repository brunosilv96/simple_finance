# Simple Finance API

API REST de finanças pessoais em Go.

## Stack atual

- Go 1.25+
- Gin (HTTP framework)
- UUID (github.com/google/uuid)
- Repositórios in-memory

## Estado atual do projeto

O projeto está em transição para uma arquitetura mais limpa, com separação entre:

- app/bootstrap (composição de dependências)
- http (router, handlers, dto)
- finance (domínio e casos de uso)
- infra (implementações de repositório)

Estrutura principal:

```text
cmd/
  api/
    main.go

internal/
  app/
    bootstrap.go
  http/
    dto/
    handlers/
    middleware/
    router/
  finance/
    category/
      entity/
      usecase/
      category_repository.go
      errors.go
    user/
      entity/
      usecase/
      user_repository.go
      errors.go
    debit/
      entity/
      usecase/
      errors.go
  infra/
    repository/
      memory_category_repository.go
      memory_user_repository.go
  shared/
    app_error.go
```

## Rotas disponíveis

### Health

| Método | Rota      | Descrição                    |
|--------|-----------|------------------------------|
| GET    | /health   | Verifica disponibilidade API |

### User

| Método | Rota              | Descrição                 |
|--------|-------------------|---------------------------|
| POST   | /api/v1/users     | Cria um novo usuário      |
| GET    | /api/v1/users/:id | Busca usuário por ID      |

### Category

| Método | Rota                   | Descrição                    |
|--------|------------------------|------------------------------|
| POST   | /api/v1/categories     | Cria uma nova categoria      |
| GET    | /api/v1/categories     | Lista todas as categorias    |
| GET    | /api/v1/categories/:id | Busca categoria por ID       |
| DELETE | /api/v1/categories/:id | Remove categoria por ID      |

Observação: o domínio debit já possui entidade e caso de uso, mas ainda não está exposto por rotas HTTP.

## Exemplos de uso

### Health

```bash
curl http://localhost:8080/health
```

### Criar usuário

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Bruno Silva"}'
```

### Buscar usuário por ID

```bash
curl http://localhost:8080/api/v1/users/<id>
```

### Criar categoria

```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{"name":"Alimentacao","description":"Mercado e restaurantes"}'
```

### Listar categorias

```bash
curl http://localhost:8080/api/v1/categories
```

### Buscar categoria por ID

```bash
curl http://localhost:8080/api/v1/categories/<id>
```

### Remover categoria

```bash
curl -X DELETE http://localhost:8080/api/v1/categories/<id>
```

## Como executar

Pré-requisito:

- Go 1.25+

Passos:

1. Instalar dependências:

```bash
go mod tidy
```

2. Subir aplicação:

```bash
go run ./cmd/api/main.go
```

Servidor padrão: http://localhost:8080

## Qualidade e limitações atuais

- Persistência in-memory (dados se perdem ao reiniciar)
- Ainda sem autenticação/autorização
- Ainda sem suíte de testes automatizados
- Ainda sem graceful shutdown explícito no main

## Próximos passos sugeridos

1. Adicionar testes de use case e handlers.
2. Adicionar graceful shutdown no servidor HTTP.
3. Expor endpoints do domínio debit.
4. Evoluir repositório para persistência real (PostgreSQL, por exemplo).
5. Padronizar respostas de erro em um formato único na camada HTTP.
