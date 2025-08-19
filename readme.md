# Go ToDo API - Clean Architecture
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
[![Postman](https://img.shields.io/badge/Postman-FF6C37?style=for-the-badge&logo=postman&logoColor=white)](https://www.postman.com/dan-3200/workspace/publico/collection/43029232-dfa83f1a-5ff2-47f7-a0ab-7cf1d0adb96c?action=share&creator=43029232)

## Descrição 

API RESTful para gerenciamento de tarefas (ToDos), desenvolvida em Golang com o framework HTTP Gin e banco de dados PostgreSQL, seguindo o padrão Clean Architecture. Permite criar, listar, atualizar e remover tarefas por meio de endpoints REST.

O projeto foi criado com foco em aprendizado, aplicando princípios de design como SOLID e o padrão Singleton.

## Tecnologias
- `Golang 1.24.1`
- `Gin (framework http)`
- `PostgreSQL`
- `Docker`

## Estrutura do projeto (Clean Architecture)
```
├── internal
│   ├── contracts       # Interfaces e definições de contratos
│   ├── controller      # Controladores (camada de entrega)
│   ├── db              # Conexão e setup de banco de dados
│   ├── dto             # Data Transfer Objects
│   ├── model           # Entidades (models/domínio)
│   ├── repository      # Implementação dos repositórios
│   ├── server          # Inicialização e rotas do servidor
│   ├── tests           # Testes automatizados
│   └── usecase         # Casos de uso (regras de negócio)
├── pkg                 # Pacotes reutilizáveis
├── main.go             # Ponto de entrada da aplicação
├── go.mod              # Dependências do Go
├── go.sum              # Hash das dependências
├── .env.example        # Exemplo de variáveis de ambiente
├── docker-compose.yml  # Configuração de containers
├── Makefile            # Atalhos de build/comando
└── README.md
```
---

### Model (Entity ToDo)
```golang
type ToDo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
```

## Instrução de Instalação 
```bash
git clone https://github.com/DAN-3200/go-todo-service.git
cd go-todo-service
go run .
```
## Instrução de uso

```bash
cd go-todo-service
make <tag>
```
---
```Makefile
# <tag>

# executar codigo
run: 
   go run .

# hot-reload
dev: 
	air

# rodar todos os tests (unitários e integração)
test: 
	go test ./...

# rodar banco principal no docker
db-main:
	docker compose -f 'docker-compose.yml' up -d --build 'database'

 # rodar banco para tests (integração) no docker
db-test:
	docker compose -f 'docker-compose.yml' up -d --build 'db-test'
```
> [!IMPORTANT]
> Para execução do docker, certifique-se de configurar o arquivo `.env` com base no `.env.example`.


## API Endpoints
| Route                   | Descrição                                  |
|-------------------------|---------------------------------------------|
| <kbd>POST /todo</kbd>     | Criar uma nova tarefa                     |
| <kbd>GET /todo/:id</kbd>  | Buscar tarefa específica por ID          |
| <kbd>GET /todo</kbd>      | Listar todas as tarefas                  |
| <kbd>PATCH /todo/:id</kbd>| Editar uma tarefa existente por ID   |
| <kbd>DELETE /todo/:id</kbd>| Remover uma tarefa existente por ID     |

> [Link publico do Postman](https://www.postman.com/dan-3200/workspace/publico/collection/43029232-dfa83f1a-5ff2-47f7-a0ab-7cf1d0adb96c?action=share&creator=43029232)

## Licença
Este projeto está licenciado sob a Licença MIT. Consulte o arquivo [LICENSE](./LICENSE) para mais detalhes.
