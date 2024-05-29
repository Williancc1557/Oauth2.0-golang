BINARY_NAME=main.out

.PHONY: all test clean

test:
	go test ./test/..._test -v