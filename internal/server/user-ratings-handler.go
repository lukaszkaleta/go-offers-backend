package server

import (
	"encoding/json"
	"fmt"
	"naborly/internal/api/rating"
	"net/http"
)

func UserRatingsRoutes(apiServer *APIServer, router *http.ServeMux) {
	router.HandleFunc("/users/{id}/ratings", makeHttpHandlerFunc(apiServer.handleUserRatings))
}

func (apiServer *APIServer) handleUserRatings(
	w http.ResponseWriter, r *http.Request) error {
	idStr, id, err := apiServer.Id(r)
	if err != nil {
		return fmt.Errorf("invalid user id: %s", idStr)
	}
	ratingPayload := new(rating.RatingModel)
	if err := json.NewDecoder(r.Body).Decode(ratingPayload); err != nil {
		return err
	}

	user, err := apiServer.Users.ById(id)
	if err != nil {
		return err
	}
	user.Ratings().Add(*ratingPayload)

	return WriteJson(w, http.StatusOK, ratingPayload)
}
