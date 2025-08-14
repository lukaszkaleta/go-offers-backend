package offer

import "naborly/internal/api/common"

type Offers interface {
	AddFromPosition(model *common.PositionModel) (Offer, error)
}

func OfferModels(offers []Offer) []*OfferModel {
	var models []*OfferModel
	for _, modelAware := range offers {
		models = append(models, modelAware.Model()) // note the = instead of :=
	}
	return models
}

func OfferHints(offers []Offer) []*OfferHint {
	var hints []*OfferHint
	for _, o := range offers {
		if o != nil {
			hints = append(hints, o.Model().Hint()) // note the = instead of :=
		}
	}
	return hints
}
