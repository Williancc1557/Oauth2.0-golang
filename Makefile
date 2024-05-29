BINARY_NAME=main.out

.PHONY: all test clean

run:
	go build
	go auth

test:
	go test ./test/..._test -v