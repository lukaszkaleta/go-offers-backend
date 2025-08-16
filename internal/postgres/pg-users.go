package postgres

import (
	"database/sql"
	"naborly/internal/api/common"
	"naborly/internal/api/user"
)

type PgUsers struct {
	DB *PgDb
}

func NewPgUsers(s *PgDb) user.Users {
	return &PgUsers{DB: s}
}

func (pgUsers *PgUsers) Add(model *common.PersonModel) (user.User, error) {
	userId := 0
	query := "INSERT INTO users(person_first_name, person_last_name, person_email, person_phone) VALUES( $1, $2, $3, $4 ) returning id"
	row := pgUsers.DB.Database.QueryRow(query, model.FirstName, model.LastName, model.Email, model.Phone)
	row.Scan(&userId)
	pgUser := PgUser{
		DB: pgUsers.DB,
		ID: userId,
	}
	return user.NewSolidUser(
		&user.UserModel{Id: userId, Person: model, Address: &common.AddressModel{}},
		pgUser,
		userId,
	), nil
}

func (pgUsers *PgUsers) ById(id int) (user.User, error) {
	personRow := new(common.PersonModel)
	addressRow := new(common.AddressModel)
	query := "select * from users where id = $1"
	row := pgUsers.DB.Database.QueryRow(query, id)
	err := row.Scan(
		&id,
		&personRow.FirstName,
		&personRow.LastName,
		&personRow.Email,
		&personRow.Phone,
		&addressRow.Line1,
		&addressRow.Line2,
		&addressRow.City,
		&addressRow.PostalCode,
		&addressRow.District,
	)
	pgUser := PgUser{DB: pgUsers.DB, ID: id}
	if err != nil {
		return pgUser, err
	}

	return user.NewSolidUser(
		&user.UserModel{Id: id, Person: personRow, Address: addressRow},
		pgUser,
		id,
	), nil
}

func (pgUsers *PgUsers) ListAll() ([]user.User, error) {
	query := "select * from users"
	rows, err := pgUsers.DB.Database.Query(query)
	if err != nil {
		return nil, err
	}

	users := []user.User{}
	for rows.Next() {
		userModel := new(user.UserModel)
		id := 0
		err := userRowScan(rows, userModel, &id)
		if err != nil {
			return nil, err
		}
		pgUser := PgUser{DB: pgUsers.DB, ID: id}
		solidUser := user.NewSolidUser(userModel, pgUser, id)
		users = append(users, solidUser)
	}
	return users, nil
}

func userRowScan(row *sql.Rows, userRow *user.UserModel, id *int) error {
	return row.Scan(
		&id,
		&userRow.Person.FirstName,
		&userRow.Person.LastName,
		&userRow.Person.Email,
		&userRow.Person.Phone,
		&userRow.Address.Line1,
		&userRow.Address.Line2,
		&userRow.Address.City,
		&userRow.Address.PostalCode,
		&userRow.Address.District,
	)
}
