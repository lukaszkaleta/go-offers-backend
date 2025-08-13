package server

import (
	"encoding/json"
	"fmt"
	"naborly/internal/api/common"
	"net/http"
)

func UserRoutes(s *APIServer, router *http.ServeMux) {
	router.HandleFunc("/users/{id}", makeHttpHandlerFunc(s.handleUser))
	router.HandleFunc("/users/{id}/address", makeHttpHandlerFunc(s.handleUserAddress))
	router.HandleFunc("/users/{id}/person", makeHttpHandlerFunc(s.handleUserPerson))
}

func (s *APIServer) handleUser(
	w http.ResponseWriter,
	r *http.Request) error {
	idStr, id, err := s.Id(r)
	if err != nil {
		return fmt.Errorf("invalid user id: %s", idStr)
	}

	if r.Method == "GET" {
		return s.handleGetUser(w, r, id)
	}
	return nil
}

func (s *APIServer) handleGetUser(
	w http.ResponseWriter, r *http.Request, id int) error {
	user, err := s.Users.ById(id)
	if err != nil {
		return err
	}
	return WriteJson(w, 200, user.Model())
}

func (s *APIServer) handleUpdateUser(
	w http.ResponseWriter, r *http.Request, id int) error {
	return nil
}

func (s *APIServer) handleUserAddress(
	w http.ResponseWriter, r *http.Request) error {
	idStr, id, err := s.Id(r)
	if err != nil {
		return fmt.Errorf("invalid user id: %s", idStr)
	}
	addressPayload := new(common.AddressModel)
	if err := json.NewDecoder(r.Body).Decode(addressPayload); err != nil {
		return err
	}

	user, err := s.Users.ById(id)
	if err != nil {
		return err
	}
	if err := user.Address().Update(addressPayload); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, user.Model())
}

func (s *APIServer) handleUserPerson(
	w http.ResponseWriter, r *http.Request) error {
	idStr, id, err := s.Id(r)
	if err != nil {
		return fmt.Errorf("invalid user id: %s", idStr)
	}
	personPayload := new(common.PersonModel)
	if err := json.NewDecoder(r.Body).Decode(personPayload); err != nil {
		return err
	}

	u, err := s.Users.ById(id)
	if err != nil {
		return err
	}

	if err := u.Person().Update(personPayload); err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, personPayload)
}
