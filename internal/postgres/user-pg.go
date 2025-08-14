package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/offer"
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
	return NewPgUserSettings(pgUser.DB, pgUser.ID)
}

func (pgUser PgUser) Archive() error {
	return nil
}

func (pgUser PgUser) Offers() offer.Offers {
	tableEntity := pgUser.tableEntity()
	return PgRelationOffers{
		DB:       pgUser.DB,
		offers:   &PgOffers{DB: pgUser.DB},
		relation: tableEntity.RelationEntityWithColumnName("user_offer", "user_id"),
	}
}

func (pgUser PgUser) tableEntity() TableEntity {
	return pgUser.DB.tableEntity("users", pgUser.ID)
}
