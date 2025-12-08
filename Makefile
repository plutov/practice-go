lint:
	go tool golangci-lint run

test:
	go test -race -v ./...
