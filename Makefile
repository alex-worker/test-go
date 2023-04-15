run: build start

build:
	go build -v src/main.go

start:
	go run src/main.go
