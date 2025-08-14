package common

// API

type Price interface {
	Update(newModel *PriceModel) error
	Model() *PriceModel
}

// Builder

func PriceFromModel(model *PriceModel) Price {
	return SolidPrice{
		model: model,
	}
}

// Model

type PriceModel struct {
	Value    int64  `json:"value"`
	Currency string `json:"currency"`
}

func (model *PriceModel) Change(newModel *PriceModel) {
	model.Value = newModel.Value
	model.Currency = newModel.Currency
}

// Solid

type SolidPrice struct {
	model *PriceModel
	Price Price
}

func NewSolidPrice(model *PriceModel, Price Price) SolidPrice {
	return SolidPrice{model, Price}
}

func (addr SolidPrice) Update(newModel *PriceModel) error {
	addr.model.Change(newModel)
	if addr.Price == nil {
		return nil
	}
	return addr.Price.Update(newModel)
}

func (addr SolidPrice) Model() *PriceModel {
	return addr.model
}
