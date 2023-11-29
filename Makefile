APP_BINARY=appBinary

run:
	go run ./cmd

build:
	go build -o ./bin/${APP_BINARY} ./cmd

up:
	docker-compose up

down:
	docker-compose down

start: up build
