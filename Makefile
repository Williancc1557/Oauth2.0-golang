BINARY_NAME=main.out

.PHONY: all test clean

run:
	go build
	go auth

test:
	go test ./test/..._test -v

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html