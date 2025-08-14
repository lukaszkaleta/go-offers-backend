package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgGlobalOffers struct {
	DB *PgDb
}

func NewPgGlobalOffers(DB *PgDb) offer.GlobalOffers {
	return PgGlobalOffers{DB}
}

func (globalOffers PgGlobalOffers) NearBy(radar *common.RadarModel) ([]offer.Offer, error) {
	offers := make([]offer.Offer, 2)
	offer1 := offer.NewSolidOffer(
		nil,
		&offer.OfferModel{
			Price:    &common.PriceModel{Value: 120000, Currency: "NOK"},
			Position: &common.PositionModel{},
		},
		1)
	offer2 := offer.NewSolidOffer(
		nil,
		&offer.OfferModel{
			Price:    &common.PriceModel{Value: 50099, Currency: "NOK"},
			Position: &common.PositionModel{},
		},
		1)
	offers = append(offers, offer1, offer2)
	return offers, nil
}
