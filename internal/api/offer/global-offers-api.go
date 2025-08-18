package offer

import "naborly/internal/api/common"

type GlobalOffers interface {
	NearBy(position *common.RadarModel) ([]Offer, error)
	ById(id int) (Offer, error)
}
