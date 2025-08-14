package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/rating"
	"naborly/internal/api/user"
)

type PgUser struct {
	DB *PgDb
	ID int
}

func (pgUser PgUser) Address() common.Address {
	return PgAddress{pgUser.DB, pgUser.tableEntity()}
}

func (pgUser PgUser) Model() *user.UserModel {
	return &user.UserModel{}
}

func (pgUser PgUser) Person() common.Person {
	return PgPerson{pgUser.DB, pgUser.tableEntity()}
}

func (pgUser PgUser) Ratings() rating.Ratings {
	return NewPgRatings(pgUser.DB, pgUser.tableEntity())
}

func (pgUser PgUser) Settings() user.UserSettings {
	return NewPgSettings(pgUser.DB)
}

func (pgUser PgUser) Archive() error {
	return nil
}

func (pgUser PgUser) tableEntity() TableEntity {
	return pgUser.namedTableEntity("business_user")
}

func (pgUser PgUser) namedTableEntity(name string) TableEntity {
	return TableEntity{Name: name, Id: pgUser.ID}
}
