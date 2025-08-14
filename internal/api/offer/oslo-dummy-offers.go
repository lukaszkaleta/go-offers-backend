package offer

import (
	"math/rand"
	"naborly/internal/api/common"
)

const osloNorth = 60017857
const osloSouth = 59917508
const osloWest = 10412946
const osloEast = 11140104
const priceMax = 1000000
const priceMin = 100000

func randomPosition() *common.PositionModel {
	lat := rand.Intn(osloNorth-osloSouth) + osloSouth
	lon := rand.Intn(osloEast-osloWest) + osloWest
	return &common.PositionModel{lat, lon}
}

func randomPrice() *common.PriceModel {
	return &common.PriceModel{
		Value:    rand.Intn(priceMax-priceMin) + priceMin,
		Currency: "NOK",
	}
}

func randomOffer(id int) OfferHint {
	return OfferHint{
		Id:       id,
		Position: randomPosition(),
		Price:    randomPrice().UserFriendly(),
	}
}

func randomOffers() []OfferHint {
	offers := []OfferHint{}
	for i := range 1000 {
		offer := randomOffer(i)
		offers = append(offers, offer)
	}
	return offers
}
