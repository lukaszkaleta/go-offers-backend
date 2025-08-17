package server

import (
	"encoding/json"
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
	"net/http"
)

func GlobalOfferRoutes(s *APIServer, router *http.ServeMux) {
	router.HandleFunc("/offers/nearby", makeHttpHandlerFunc(s.handleGlobalOffersNearBy))
	router.HandleFunc("/offers/nearby/forMap", makeHttpHandlerFunc(s.handleGlobalOffersNearByForMap))
	router.HandleFunc("/offers/preview/{offerId}", makeHttpHandlerFunc(s.handleGlobalOfferPreview))
}

func (s *APIServer) handleGlobalOffersNearBy(w http.ResponseWriter, r *http.Request) error {
	radarPayload := new(common.RadarModel)
	if err := json.NewDecoder(r.Body).Decode(radarPayload); err != nil {
		return err
	}
	offers, err := s.GlobalOffers.NearBy(radarPayload)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, offer.OfferModels(offers))
}

func (s *APIServer) handleGlobalOffersNearByForMap(w http.ResponseWriter, r *http.Request) error {
	radarPayload := new(common.RadarModel)
	if err := json.NewDecoder(r.Body).Decode(radarPayload); err != nil {
		return err
	}
	offers, err := s.GlobalOffers.NearBy(radarPayload)
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, offer.OfferHints(offers))
}

func (s *APIServer) handleGlobalOfferPreview(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, nil)
}

func (s *APIServer) handleVipsCallback(w http.ResponseWriter, r *http.Request) error {
	return WriteJson(w, http.StatusOK, nil)
}
