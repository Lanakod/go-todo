include .env

build:
	go build -o .bin/main cmd/main.go

prod: build
	./.bin/main

dev:
	go run cmd/main.go

migrate_up:
	migrate -path ./schema -database 'postgres://lanakod:password@10.10.22.7:5432/go-todo?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://lanakod:password@10.10.22.7:5432/go-todo?sslmode=disable' down

swag:
	swag init -g ./cmd/main.go