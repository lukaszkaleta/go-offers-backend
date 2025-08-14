package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
)

type PgOffer struct {
	DB *PgDb
	ID int
}

func (p PgOffer) Model() *offer.OfferModel {
	//TODO implement me
	panic("implement me")
}

func (p PgOffer) Address() common.Address {
	return PgAddress{p.DB, p.tableEntity()}
}

func (p PgOffer) Position() common.Position {
	return PgPosition{p.DB, p.tableEntity()}
}

func (p PgOffer) Price() common.Price {
	return PgPrice{p.DB, p.tableEntity()}
}

func (p PgOffer) tableEntity() TableEntity {
	return p.DB.tableEntity("offer", p.ID)
}
