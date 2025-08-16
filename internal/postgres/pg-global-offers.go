package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgGlobalOffers struct {
	DB *PgDb
}

func NewPgGlobalOffers(DB *PgDb) offer.GlobalOffers {
	return &PgGlobalOffers{DB}
}

func (globalOffers *PgGlobalOffers) NearBy(radar *common.RadarModel) (*[]offer.Offer, error) {
	id := 0
	query := "select * from offer"
	offers := []offer.Offer{}
	rows, err := globalOffers.DB.Database.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		addressRow := new(common.AddressModel)
		priceRow := new(common.PriceModel)
		positionRow := new(common.PositionModel)
		err := rows.Scan(
			&id,
			&addressRow.Line1,
			&addressRow.Line2,
			&addressRow.City,
			&addressRow.PostalCode,
			&addressRow.District,
			&positionRow.Lat,
			&positionRow.Lon,
			&priceRow.Value,
			&priceRow.Currency,
		)
		pgOffer := &PgOffer{DB: globalOffers.DB, ID: id}
		if err != nil {
			return nil, err
		}

		solidOffer := offer.NewSolidOffer(
			&offer.OfferModel{
				Id:       id,
				Address:  addressRow,
				Price:    priceRow,
				Position: positionRow,
			},
			pgOffer,
			id)
		offers = append(offers, solidOffer)
	}
	return &offers, nil
}
