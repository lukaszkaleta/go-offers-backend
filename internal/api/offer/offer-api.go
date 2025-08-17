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
	Description() common.Description
}

// Model

type OfferModel struct {
	Id          int                      `json:"id"`
	Position    *common.PositionModel    `json:"position"`
	Price       *common.PriceModel       `json:"price"`
	Address     *common.AddressModel     `json:"address"`
	Description *common.DescriptionModel `json:"description"`
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
	return &SolidOffer{
		id,
		model,
		Offer,
	}
}
func (solidOffer *SolidOffer) Model() *OfferModel {
	return solidOffer.model
}

func (solidOffer *SolidOffer) Position() common.Position {
	if solidOffer.Offer != nil {
		return common.NewSolidPosition(
			solidOffer.Model().Position,
			solidOffer.Offer.Position(),
		)
	}
	return common.NewSolidPosition(solidOffer.Model().Position, nil)
}

func (solidOffer *SolidOffer) Price() common.Price {
	if solidOffer.Offer != nil {
		return common.NewSolidPrice(
			solidOffer.Model().Price,
			solidOffer.Offer.Price(),
		)
	}
	return common.NewSolidPrice(solidOffer.Model().Price, nil)
}

func (solidOffer *SolidOffer) Address() common.Address {
	if solidOffer.Offer != nil {
		return common.NewSolidAddress(
			solidOffer.Model().Address,
			solidOffer.Offer.Address(),
		)
	}
	return common.NewSolidAddress(solidOffer.Model().Address, nil)
}
func (solidOffer *SolidOffer) Description() common.Description {
	if solidOffer.Offer != nil {
		return common.NewSolidDescription(
			solidOffer.Model().Description,
			solidOffer.Offer.Description(),
		)
	}
	return common.NewSolidDescription(solidOffer.Model().Description, nil)
}
