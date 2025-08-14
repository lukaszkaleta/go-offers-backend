package server

import (
	"encoding/json"
	"fmt"
	"log"
	"naborly/internal/api/offer"
	"naborly/internal/api/user"
	"naborly/internal/postgres"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

func WriteJson(w http.ResponseWriter, status int, payload any) error {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHttpHandlerFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr   string
	Users        user.Users
	GlobalOffers offer.GlobalOffers
}

func NewAPIServer(listenAddr string, pgDb *postgres.PgDb) *APIServer {
	users := postgres.NewPgUsers(pgDb)
	offers := postgres.NewPgGlobalOffers(pgDb)
	return &APIServer{
		listenAddr:   listenAddr,
		Users:        users,
		GlobalOffers: offers,
	}
}

func (apiServer *APIServer) Run() {
	router := http.NewServeMux()
	UsersRoutes(apiServer, router)
	UserRoutes(apiServer, router)
	UserRatingsRoutes(apiServer, router)
	UserRatingRoutes(apiServer, router)
	GlobalOfferRoutes(apiServer, router)

	log.Println("Api Server listening on", apiServer.listenAddr)
	http.ListenAndServe(apiServer.listenAddr, router)
}

func (s *APIServer) Id(r *http.Request) (string, int, error) {
	return s.NamedId("id", r)
}

func (s *APIServer) NamedId(name string, r *http.Request) (string, int, error) {
	idStr := r.PathValue(name)
	atoi, err := strconv.Atoi(idStr)
	return idStr, atoi, err
}

// JWT

func withJwtAuth(handlerFunc http.HandlerFunc, users user.Users) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("calling JWT auth midleware\n")
		tokenString := r.Header.Get("Authorization")
		claims := validateJwt(tokenString)
		println(claims)
	}
}

type naborlyClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func validateJwt(tokenString string) *naborlyClaims {
	token, err := jwt.ParseWithClaims(tokenString, &naborlyClaims{}, func(token *jwt.Token) (any, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		permissionDenied(nil)
		return nil
	}
	claims, ok := token.Claims.(*naborlyClaims)
	if ok {
		fmt.Println(claims.UserId, claims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
		return nil
	}

	return claims
}

func permissionDenied(w http.ResponseWriter) {
	WriteJson(w, http.StatusForbidden, ApiError{Error: "permission denied"})
}
