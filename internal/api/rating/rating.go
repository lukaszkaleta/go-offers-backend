package rating

// API
type Rating interface {
	Model() *RatingModel
	Update(rating Rating) error
}

// Model

type RatingModel struct {
	Description string
	Score       int
}

func (m RatingModel) Change(rating Rating) {
	m.Score = rating.Model().Score
	m.Description = rating.Model().Description
}

// Solid

type SolidRating struct {
	Id     int
	model  *RatingModel
	rating Rating
}

func NewSolidRating(ratingModel *RatingModel, rating Rating, id int) SolidRating {
	return SolidRating{
		model:  ratingModel,
		rating: rating,
		Id:     id,
	}
}

func (s SolidRating) Model() *RatingModel {
	return s.model
}

func (s SolidRating) Update(rating Rating) error {
	s.model.Change(rating)
	return nil
}
