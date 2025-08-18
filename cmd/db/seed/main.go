package main

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
	"naborly/internal/api/user"
	"naborly/internal/postgres"
)

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

func SeedOffers(user user.User) error {
	offers := user.Offers()
	positions := offer.RandomPositions(10)
	for _, position := range positions {
		o, err := offers.AddFromPosition(position)
		if err != nil {
			return err
		}
		err = o.Price().Update(offer.RandomPrice())
		if err != nil {
			return err
		}
		err = o.Description().Update(offer.RandomDescription())
		if err != nil {
			return err
		}
	}
	return nil
}

func SeedUsers(users user.Users) ([]user.User, error) {
	data := []user.User{}
	lukasz, err := users.Add(
		&common.PersonModel{
			Phone:     "004794004108",
			FirstName: "Lukasz",
			LastName:  "Kaleta",
			Email:     "lukaszkaleta@gmail.com",
		},
	)
	if err != nil {
		return nil, err
	}
	oystein, err := users.Add(
		&common.PersonModel{
			Phone:     "0047",
			FirstName: "Oystein",
			LastName:  "Bakke",
			Email:     "",
		},
	)
	if err != nil {
		return nil, err
	}
	data = append(data, lukasz, oystein)
	return data, nil
}
