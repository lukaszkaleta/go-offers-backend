package logging

import (
	"fmt"
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
	"time"
)

type GlobalOffersLogger struct {
	next offer.GlobalOffers
}

func NewGlobalOffersLogger(next offer.GlobalOffers) offer.GlobalOffers {
	return &GlobalOffersLogger{
		next: next,
	}
}

func (g *GlobalOffersLogger) NearBy(position *common.RadarModel) (*[]offer.Offer, error) {
	defer func(start time.Time) {
		fmt.Printf("Searching for offers took %v\n", time.Since(start))
	}(time.Now())
	return g.next.NearBy(position)
}
