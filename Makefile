generate: generated.go

generated.go: .PHONY
	go get github.com/99designs/gqlgen/cmd@0.13.0
	go run github.com/99designs/gqlgen generate