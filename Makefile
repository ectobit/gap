.PHONY: test test-cov

lint:
	@golangci-lint run

test:
	@go test -race ./...

test-cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out
