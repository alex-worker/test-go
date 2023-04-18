run: build start

build:
	go build -v src/main.go

start:
	go run src/main.go

test:
	go test ./src/core/* -v

update:
	go get -u all
	go mod tidy
