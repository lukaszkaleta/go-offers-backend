package postgres

import (
	"naborly/internal/api/rating"
)

type PgRating struct {
	db *PgDb
	Id int
}

func NewPgRating(db *PgDb, id int) rating.Rating {
	return &PgRating{db, id}
}

func (pgRating *PgRating) Model() *rating.RatingModel {
	return &rating.RatingModel{}
}

func (pgRating *PgRating) Update(newModel *rating.RatingModel) error {
	return nil
}
