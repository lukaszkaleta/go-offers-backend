package server

import (
	"encoding/json"
	"fmt"
	"naborly/internal/api/rating"
	"net/http"
)

func UserRatingRoutes(apiServer *APIServer, router *http.ServeMux) {
	router.HandleFunc("/users/{id}/ratings/{ratingId}", makeHttpHandlerFunc(apiServer.handleUserRating))
}

func (apiServer *APIServer) handleUserRating(
	w http.ResponseWriter, r *http.Request) error {
	idStr, id, err := apiServer.Id(r)
	if err != nil {
		return fmt.Errorf("invalid user id: %s", idStr)
	}
	ratingIdStr, ratingId, err := apiServer.NamedId("ratingId", r)
	if err != nil {
		return fmt.Errorf("invalid user rating id: %s", ratingIdStr)
	}

	newModel := new(rating.RatingModel)
	if err := json.NewDecoder(r.Body).Decode(newModel); err != nil {
		return err
	}

	user, err := apiServer.Users.ById(id)
	if err != nil {
		return err
	}
	userRating, err := user.Ratings().ById(ratingId)
	if err != nil {

	}
	userRating.Update(newModel)

	return WriteJson(w, http.StatusOK, newModel)
}
