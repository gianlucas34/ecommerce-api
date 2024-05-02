.PHONY: run
run:
	go mod tidy
	go run ./cmd/server/main.go

.PHONY: unit-test
unit-test:
	go test -v ./internal/...
