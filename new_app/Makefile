# export PATH=$PATH:~/go/go1.22.2/bin
# export PATH=$PATH:$GOPATH/bin
export CGO_ENABLED=1

run: build start

build:
	go build -v src/main.go

start:
	go run src/main.go

test:
	go test ./src/math/... -v

update:
	go get -u all
	go mod tidy
