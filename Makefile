# Makefile to build appconfiguration-admin-sdk-go

build:
	go build ./...

runUnitTests:
	make build
	cd appconfigurationv1 && go test
	cd common && go test

tidy:
	go mod tidy
