package offer

import (
	"naborly/internal/api/common"
)

// API

type Offer interface {
	Model() *OfferModel
	Address() common.Address
	Position() common.Position
}

// Model

type OfferModel struct {
	Id       int                   `json:"id"`
	Position *common.PositionModel `json:"position"`
	Address  *common.AddressModel  `json:"address"`
}

// Builder

func NewSolidOffer(Offer Offer, model *OfferModel, id int) Offer {
	return SolidOffer{
		id,
		model,
		Offer,
	}
}

// Solid

type SolidOffer struct {
	Id    int
	model *OfferModel
	Offer Offer
}

func (u SolidOffer) Model() *OfferModel {
	return u.model
}

func (u SolidOffer) Position() common.Position {
	if u.Offer != nil {
		return common.NewSolidPosition(
			u.Model().Position,
			u.Offer.Position(),
		)
	}
	return common.NewSolidPosition(u.Model().Position, nil)
}

func (u SolidOffer) Address() common.Address {
	if u.Offer != nil {
		return common.NewSolidAddress(
			u.Model().Address,
			u.Offer.Address(),
		)
	}
	return common.NewSolidAddress(u.Model().Address, nil)
}
