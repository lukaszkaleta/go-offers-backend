package main

import "naborly/internal/postgres"

func main() {
	pg := postgres.NewPg()
	users, err := SeedUsers(postgres.NewPgUsers(pg))
	panicCheck(err)
	for _, user := range users {
		err := SeedOffers(user)
		panicCheck(err)
	}
}

func panicCheck(err error) {
	if err != nil {
		panic(err)
	}
}
