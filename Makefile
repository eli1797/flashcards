generate: generated.go

generated.go: .PHONY
	go get github.com/99designs/gqlgen/cmd@v0.13.0
	go run github.com/99designs/gqlgen generate

precommit:
	go get -u ./...
	go mod tidy
	go test ./...
	git add go.mod
	git add go.sum

test:
	go test ./...
# 	go install github.com/jstemmer/go-junit-report

build:
	go get ./..
	go build -o bin/main main.go

build-linux:
	go get ./...
	export GIN_MODE=release 
	export GOOS=linux 
	go build -o bin/main main.go

run:
	go run main.go
	
.PHONY: