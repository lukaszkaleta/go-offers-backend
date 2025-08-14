package offer

import (
	"naborly/internal/api/common"
)

// API

type Offer interface {
	Model() *OfferModel
	Address() common.Address
	Position() common.Position
	Price() common.Price
}

// Model

type OfferModel struct {
	Id       int                   `json:"id"`
	Position *common.PositionModel `json:"position"`
	Price    *common.PriceModel    `json:"price"`
	Address  *common.AddressModel  `json:"address"`
}

func (m *OfferModel) Hint() *OfferHint {
	return &OfferHint{
		Position: m.Position,
		Price:    m.Price.UserFriendly(),
	}
}

type OfferHint struct {
	Id       int                   `json:"id"`
	Position *common.PositionModel `json:"position"`
	Price    string                `json:"price"`
}

// Solid

type SolidOffer struct {
	Id    int
	model *OfferModel
	Offer Offer
}

func NewSolidOffer(model *OfferModel, Offer Offer, id int) Offer {
	return SolidOffer{
		id,
		model,
		Offer,
	}
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

func (u SolidOffer) Price() common.Price {
	if u.Offer != nil {
		return common.NewSolidPrice(
			u.Model().Price,
			u.Offer.Price(),
		)
	}
	return common.NewSolidPrice(u.Model().Price, nil)
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
