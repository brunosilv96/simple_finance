# Simple Finance API

API REST simples de finanças pessoais desenvolvida em Go puro, sem frameworks externos. Permite gerenciar usuários, categorias e débitos financeiros.

## Tecnologias

- **Go** 1.25+
- **net/http** (stdlib) — servidor HTTP
- **github.com/google/uuid** — geração de IDs únicos

## Arquitetura

O projeto segue os princípios de **Clean Architecture**, separando responsabilidades em camadas:

```
cmd/
└── main.go                  # Ponto de entrada e injeção de dependências

internal/
└── finance/
    ├── dto/                 # Data Transfer Objects (request/response)
    ├── entity/              # Entidades de domínio
    ├── handler/             # Handlers HTTP (controllers)
    ├── repository/          # Implementações de repositório (in-memory)
    └── usecase/             # Casos de uso (regras de negócio)
        ├── category/
        ├── debit/
        └── user/
middleware/
└── logger.go                # Middleware de logging de requisições
```

## Endpoints

### Health Check
| Método | Rota              | Descrição            |
|--------|-------------------|----------------------|
| GET    | `/api/v1/health`  | Verifica status da API |

### Usuários
| Método | Rota             | Descrição             |
|--------|------------------|-----------------------|
| POST   | `/api/v1/users`  | Registrar novo usuário |

### Categorias
| Método | Rota                      | Descrição                  |
|--------|---------------------------|----------------------------|
| POST   | `/api/v1/categories`      | Criar nova categoria        |
| GET    | `/api/v1/categories`      | Listar todas as categorias  |
| GET    | `/api/v1/categories/{id}` | Buscar categoria por ID     |
| DELETE | `/api/v1/categories/{id}` | Deletar categoria           |

## Exemplos de Uso

### Health Check
```bash
curl http://localhost:8080/api/v1/health
```
```json
{"status": "ok"}
```

### Registrar Usuário
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Bruno Silva"}'
```
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440000",
  "name": "Bruno Silva",
  "created_at": "2026-04-15T10:00:00Z"
}
```

### Criar Categoria
```bash
curl -X POST http://localhost:8080/api/v1/categories \
  -H "Content-Type: application/json" \
  -d '{"name": "Alimentação", "description": "Gastos com comida e restaurantes"}'
```
```json
{
  "id": "550e8400-e29b-41d4-a716-446655440001",
  "name": "Alimentação",
  "description": "Gastos com comida e restaurantes",
  "created_at": "2026-04-15T10:00:00Z"
}
```

### Listar Categorias
```bash
curl http://localhost:8080/api/v1/categories
```

### Buscar Categoria por ID
```bash
curl http://localhost:8080/api/v1/categories/550e8400-e29b-41d4-a716-446655440001
```

### Deletar Categoria
```bash
curl -X DELETE http://localhost:8080/api/v1/categories/550e8400-e29b-41d4-a716-446655440001
```

## Como Executar

### Pré-requisitos
- Go 1.25 ou superior instalado

### Passos

1. Clone o repositório:
```bash
git clone https://github.com/brunosilv96/simple_finance_api.git
cd simple_finance_api
```

2. Instale as dependências:
```bash
go mod tidy
```

3. Execute a aplicação:
```bash
go run ./cmd/main.go
```

O servidor estará disponível em `http://localhost:8080`.

## Observações

- O armazenamento é **in-memory**: os dados são perdidos ao reiniciar a aplicação.
- Não há autenticação implementada nesta versão.
