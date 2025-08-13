package common

// API

type Position interface {
	Update(newModel *PositionModel) error
	Model() *PositionModel
}

// Builder

// Model

type PositionModel struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"long"`
}

func (model *PositionModel) Change(newModel *PositionModel) {
	model.Lat = newModel.Lat
	model.Lon = newModel.Lon
}

// Solid

type SolidPosition struct {
	model    *PositionModel
	Position Position
}

func NewSolidPosition(model *PositionModel, Position Position) SolidPosition {
	return SolidPosition{model, Position}
}

func (addr SolidPosition) Update(newModel *PositionModel) error {
	addr.model.Change(newModel)
	if addr.Position == nil {
		return nil
	}
	return addr.Position.Update(newModel)
}

func (addr SolidPosition) Model() *PositionModel {
	return addr.model
}
