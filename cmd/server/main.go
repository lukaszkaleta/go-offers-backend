package main

import (
	"fmt"
	"naborly/internal/postgres"
	"naborly/internal/server"
)

func main() {
	pg := postgres.NewPg()
	pg.Init()
	fmt.Printf("%+v\n", pg)

	server := server.NewAPIServer(":3000", pg)
	server.Run()
}
