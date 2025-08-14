build:
	@go build -o bin/naborly-server cmd/server/main.go

## Application

run: build
	@./bin/naborly-server

test:
	@go test -v/...

## Database

drop:
	@go run cmd/db/drop/main.go

init:
	@go run cmd/db/init/main.go

seed:
	@go run cmd/db/seed/main.go
