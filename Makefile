# Makefile to build appconfiguration-go-admin-sdk
GO=go
LINT=golangci-lint
GOSEC=gosec
TEST_TAGS=
COVERAGE = -coverprofile=coverage.txt -covermode=atomic

all: tidy test lint
travis-ci: tidy test-cov lint scan-gosec

test:
	${GO} test ./... ${TEST_TAGS}

test-cov:
	${GO} test ./... ${TEST_TAGS} ${COVERAGE}

test-int:
	${GO} test ./... -tags=integration

test-int-cov:
	${GO} test ./... -tags=integration ${COVERAGE}

lint:
	${LINT} run --build-tags=integration,examples --timeout 3m

scan-gosec:
	${GOSEC} ./...

tidy:
	${GO} mod tidy
