# !windows powershell: 'choco install make' 

# Evitar conflito com nome de arquivos
.PHONY: dev test db-main db-test run

# executar código
run: 
	go run .

# hot-reload
dev: 
	air

# rodar todos os testes (unitários e integração)
test: 
	go test ./...

# rodar banco principal no docker
db-main:
	docker compose -f 'docker-compose.yml' up -d --build 'database'

# rodar banco para tests (integração) no docker
db-test:
	docker compose -f 'docker-compose.yml' up -d --build 'db-test'
