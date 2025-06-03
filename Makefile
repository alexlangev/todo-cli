.DEFAULT_GOAL := build

.PHONY:fmt vet build clean test
fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build -o todo ./cmd/todo

clean:
	go clean
	rm -f .todo.json ./todo

test:
	go test -v ./...
