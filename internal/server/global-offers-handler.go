package server

import (
	"naborly/internal/api/common"
	"net/http"
)

func GlobalOfferRoutes(s *APIServer, router *http.ServeMux) {
	router.HandleFunc("/offers/nearby/{lat}/{lon}", makeHttpHandlerFunc(s.handleGlobalOffersNearBy))
}

func (s *APIServer) handleGlobalOffersNearBy(w http.ResponseWriter, r *http.Request) error {
	lat := r.URL.Query().Get("lat")
	lon := r.URL.Query().Get("lon")
	s.GlobalOffers.NearBy(common.NewPosition(lat, lon))
	return nil
}
