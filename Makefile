# Makefile to build appconfiguration-go-admin-sdk

build:
	go build ./...

lint:
	golangci-lint run

runUnitTests:
	make build
	cd appconfigurationv1 && go test
	cd common && go test

tidy:
	go mod tidy

vendor:
	go mod vendor

runCoverage:
	make build
	cd appconfigurationv1 && go test -coverprofile=coverage.out
	cd common && go test -coverprofile=coverage.out