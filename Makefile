.PHONY: run
run:
	go run cmd/overseer/main.go

.PHONY: test-cover
test-cover:
	mkdir -p test
	go test -coverprofile=test/coverage.out ./...
	go tool cover -func=test/coverage.out

.PHONY: test
test:
	go test ./...

.PHONY: build
build:
	  go build -o bin/main ./cmd/overseer/main.go