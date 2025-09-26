BINARY_NAME=main.out

# Use tabs for the recipe lines. "make" requires tabs, not spaces.
.PHONY: build run clean test

build:
	go build -o ${BINARY_NAME} cmd/main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}

test:
	go test -v ./...
