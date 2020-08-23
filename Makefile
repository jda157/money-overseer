.PHONY: run
run:
	go run cmd/overseer/main.go

.PHONY: test-cover
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

.PHONY: test
test:
	go test ./...