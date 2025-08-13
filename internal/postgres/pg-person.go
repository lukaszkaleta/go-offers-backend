package postgres

import (
	"fmt"
	"naborly/internal/api/common"
)

type PgPerson struct {
	DB          *PgDb
	TableEntity TableEntity
}

func (pg PgPerson) Update(model *common.PersonModel) error {
	query := fmt.Sprintf("update %s set person_first_name = $1, person_last_name = $2, person_email = $3, person_phone = $4 where id = $5", pg.TableEntity.Name)
	_, err := pg.DB.Database.Exec(query, model.FirstName, model.LastName, model.Email, model.Phone, pg.TableEntity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pg PgPerson) Model() *common.PersonModel {
	return &common.PersonModel{}
}
