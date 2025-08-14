package offer

import "naborly/internal/api/common"

type GlobalOffers interface {
	NearBy(position common.Radar) ([]Offer, error)
}
