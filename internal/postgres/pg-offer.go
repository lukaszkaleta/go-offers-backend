package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgOffer struct {
	DB *PgDb
	ID int
}

func (pgOffer *PgOffer) Model() *offer.OfferModel {
	//TODO implement me
	panic("implement me")
}

func (pgOffer *PgOffer) Address() common.Address {
	return &PgAddress{pgOffer.DB, pgOffer.tableEntity()}
}

func (pgOffer *PgOffer) Position() common.Position {
	return &PgPosition{pgOffer.DB, pgOffer.tableEntity()}
}

func (pgOffer *PgOffer) Price() common.Price {
	return &PgPrice{pgOffer.DB, pgOffer.tableEntity()}
}

func (pgOffer *PgOffer) tableEntity() TableEntity {
	return pgOffer.DB.tableEntity("offer", pgOffer.ID)
}
