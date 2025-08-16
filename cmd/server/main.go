package main

import (
	"fmt"
	"naborly/internal/api/offer"
	"naborly/internal/postgres"
	"naborly/internal/server"
)

func main() {
	offer.OsloMiddle()
	pg := postgres.NewPg()
	fmt.Printf("%+v\n", pg)

	server := server.NewAPIServer(":3000", pg)
	server.Run()
}
