package postgres

import (
	"fmt"
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgOffers struct {
	DB  *PgDb
	Ids []int
}

func (pgOffers *PgOffers) AddFromPosition(model *common.PositionModel) (offer.Offer, error) {
	offerId := 0
	query := "INSERT INTO offer(position_latitude, position_longitude) VALUES( $1, $2 ) returning id"
	row := pgOffers.DB.Database.QueryRow(query, model.Lat, model.Lon)
	row.Scan(&offerId)
	pgOffer := PgOffer{
		DB: pgOffers.DB,
		ID: offerId,
	}
	return offer.NewSolidOffer(
		&offer.OfferModel{
			Id:       offerId,
			Position: model,
			Address:  &common.AddressModel{},
			Price:    &common.PriceModel{},
		},
		&pgOffer,
		offerId,
	), nil
}

// Relation

type PgRelationOffers struct {
	DB       *PgDb
	offers   *PgOffers
	relation RelationEntity
}

func NewPgRelationOffers(pfOffers *PgOffers, relation RelationEntity) PgRelationOffers {
	return PgRelationOffers{
		DB:       pfOffers.DB,
		relation: relation,
		offers:   pfOffers,
	}
}

func (p PgRelationOffers) AddFromPosition(model *common.PositionModel) (offer.Offer, error) {
	newOffer, err := p.offers.AddFromPosition(model)
	if err != nil {
		return newOffer, err
	}
	query := fmt.Sprintf("INSERT INTO %s(offer_id, %s) VALUES( $1, $2 )", p.relation.TableName, p.relation.ColumnName)
	_, err = p.DB.Database.Exec(query, newOffer.Model().Id, p.relation.RelationId)
	if err != nil {
		return newOffer, err
	}
	return newOffer, nil
}
