package postgres

import (
	"naborly/internal/api/rating"
)

type PgRating struct {
	db *PgDb
	Id int
}

func NewPgRating(db *PgDb, id int) PgRating {
	return PgRating{db, id}
}

func (pgRating PgRating) Update(rating rating.Rating) {

}
