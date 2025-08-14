package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgGlobalOffers struct {
	DB *PgDb
}

func (globalOffers PgGlobalOffers) NearBy(position common.Radar) ([]offer.Offer, error) {
	return nil, nil
}
