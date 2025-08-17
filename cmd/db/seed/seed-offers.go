package main

import (
	"naborly/internal/api/offer"
	"naborly/internal/api/user"
)

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
