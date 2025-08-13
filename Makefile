build:
	@go build -o bin/naborly-server cmd/server/main.go

run: build
	@./bin/naborly-server

test:
	@go test -v/...

drop:
	@go run cmd/db/drop/main.go

init:
	@go run cmd/db/init/main.go
