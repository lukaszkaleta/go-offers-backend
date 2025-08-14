package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
	"naborly/internal/api/user"
)

type PgOffers struct {
	DB *PgDb
}

func (pgOffers PgOffers) AddFromPosition(model *common.PositionModel) (offer.Offer, error) {
	offerId := 0
	query := "INSERT INTO offer(position_latitude, position_longitude) VALUES( $1, $2 ) returning id"
	row := pgOffers.DB.Database.QueryRow(query, model.Lat, model.Lon)
	row.Scan(&offerId)
	pgOffer := PgOffer{
		DB: pgOffers.DB,
		ID: offerId,
	}
	return offer.NewSolidOffer(
		pgOffer,
		&offer.OfferModel{
			Id:       offerId,
			Position: model,
			Address:  &common.AddressModel{},
			Price:    &common.PriceModel{},
		},
		offerId,
	), nil
}
