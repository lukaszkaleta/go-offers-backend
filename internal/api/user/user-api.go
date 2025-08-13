package user

import (
	"naborly/internal/api/common"
	"naborly/internal/api/rating"
)

// API

type User interface {
	Model() *UserModel
	Person() common.Person
	Address() common.Address
	Ratings() rating.Ratings
	Archive() error
}

// Model

type UserModel struct {
	Id      int                  `json:"id"`
	Person  *common.PersonModel  `json:"person"`
	Address *common.AddressModel `json:"address"`
}

// Builder

func NewSolidUser(user User, model *UserModel, id int) User {
	return SolidUser{
		id,
		model,
		user,
	}
}

// Solid

type SolidUser struct {
	Id    int
	model *UserModel
	user  User
}

func (u SolidUser) Model() *UserModel {
	return u.model
}

func (u SolidUser) Person() common.Person {
	if u.user != nil {
		return common.NewSolidPerson(
			u.Model().Person,
			u.user.Person(),
		)
	}
	return common.NewSolidPerson(u.Model().Person, nil)
}

func (u SolidUser) Address() common.Address {
	if u.user != nil {
		return common.NewSolidAddress(
			u.Model().Address,
			u.user.Address(),
		)
	}
	return common.NewSolidAddress(u.Model().Address, nil)
}

func (u SolidUser) Ratings() rating.Ratings {
	return u.user.Ratings()
}

func (u SolidUser) Archive() error {
	return nil
}
