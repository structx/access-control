
build:
	docker build -t trevatk/anastasia:v0.0.1 .

deps:
	go mod tidy
	go mod vendor

lint:
	golangci-lint run ./...

utest:
	go test ./...