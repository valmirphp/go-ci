# NOTES


## DOCKER RUN 
> Permite rodar um test a partir de uma imagem sem precisar criar um container

Command: `docker run --rm -v "$PWD/src":/go/src/app -w /go/src/app golang:1.14-alpine go test -v ./...`

## DOCKER-COMPOSE
> Para fins de testes, criei um docker-compose.yaml para simular executando comandos dentro de um container do compose.

Command: `docker-compose run work go test -v ./...`