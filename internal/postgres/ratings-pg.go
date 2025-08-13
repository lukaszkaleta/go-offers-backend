package postgres

import (
	"naborly/internal/api/rating"
)

type PgRatings struct {
	db         *PgDb
	ownerTable TableEntity
}

func NewPgRatings(db *PgDb, ownerTable TableEntity) PgRatings {
	return PgRatings{db: db, ownerTable: ownerTable}
}

func (s PgRatings) Add(r rating.RatingModel) (rating.Rating, error) {
	return rating.SolidRating{}, nil
}

func (s PgRatings) ById(id int) (rating.Rating, error) {
	return rating.SolidRating{}, nil
}
