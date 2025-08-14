package main

import (
	"naborly/internal/api/common"
	"naborly/internal/api/user"
)

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
