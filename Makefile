.PHONY: test test-cov

test:
	@go test -race ./...

test-cov:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -func coverage.out
