# !windows powershell: 'choco install make' 

# Evitar conflito com nome de arquivos
.PHONY: dev test db-main db-test

dev: # hot-reload
	air
test: 
	go test ./...
db-main: # banco principal
	docker compose -f 'docker-compose.yml' up -d --build 'database'
db-test: # banco secundario
	docker compose -f 'docker-compose.yml' up -d --build 'db-test'

