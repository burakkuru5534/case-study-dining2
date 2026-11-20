BIN := bin/api

.PHONY: run build test race cover vet fmt check

run:
	go run ./cmd/api

build:
	go build -o $(BIN) ./cmd/api

test:
	go test ./...

race:
	go test -race -count=1 ./...

cover:
	go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out

vet:
	go vet ./...

fmt:
	gofmt -s -w .

check: fmt vet race
