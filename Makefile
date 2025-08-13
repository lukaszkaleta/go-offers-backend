build:
	@go build -o bin/naborly-server cmd/server/main.go

run: build
	@./bin/naborly-server

test:
	@go test -v/...
	