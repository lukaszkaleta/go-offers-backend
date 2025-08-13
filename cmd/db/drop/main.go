package main

import (
	"naborly/internal/postgres"
)

func main() {
	postgres.ExecuteFromFile("cmd/db/drop/drop.sql")
}
