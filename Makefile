BINARY_NAME=main.out

.PHONY: build
build:
	go build -o ${BINARY_NAME} cmd/main.go

.PHONY: run
run:
	go build -o ${BINARY_NAME} cmd/main.go
	./${BINARY_NAME}

.PHONY: clean
clean:
	go clean
	rm ${BINARY_NAME}


.PHONY: build run clean test
test:
	go test -v ./...
