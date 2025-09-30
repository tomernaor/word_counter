all: mod vet race test build

mod:
	go mod tidy

vet:
	go vet ./...

race:
	go test -race ./...
test:
	go test -v ./...

build:
	go build -o word_counter .

