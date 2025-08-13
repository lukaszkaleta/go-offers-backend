package user

import (
	"naborly/internal/api/common"
)

type Users interface {
	Add(model common.PersonModel) (User, error)
	ById(id int) (User, error)
	ListAll() ([]User, error)
}

func UserModels(users []User) []*UserModel {
	var models []*UserModel
	for _, u := range users {
		models = append(models, u.Model()) // note the = instead of :=
	}
	return models
}
