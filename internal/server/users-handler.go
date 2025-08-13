package server

import (
	"encoding/json"
	"naborly/internal/api/common"
	"naborly/internal/api/user"
	"net/http"
)

func UsersRoutes(s *APIServer, router *http.ServeMux) {
	router.HandleFunc("/users", makeHttpHandlerFunc(s.handleUsers))
}

func (s *APIServer) handleUsers(
	w http.ResponseWriter,
	r *http.Request) error {
	if r.Method == "GET" {
		return s.handleListUsers(w, r)
	}
	if r.Method == "POST" {
		return s.handleAddUser(w, r)
	}
	return nil
}

func (s *APIServer) handleListUsers(
	w http.ResponseWriter, r *http.Request) error {
	all, err := s.Users.ListAll()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, user.UserModels(all))
}

func (apiServer *APIServer) handleAddUser(
	w http.ResponseWriter,
	r *http.Request) error {
	personPayload := new(common.PersonModel)
	if err := json.NewDecoder(r.Body).Decode(personPayload); err != nil {
		return err
	}
	user, err := apiServer.Users.Add(personPayload)
	if err != nil {
		return err
	}
	return WriteJson(w, 200, user.Model())
}
