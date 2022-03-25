SERVICE ?= portal-client

dep:
	go mod download

run:
	go run main.go

build:
	go build -o /${SERVICE} main.go