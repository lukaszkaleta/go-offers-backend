package postgres

import (
	"fmt"
	"naborly/internal/api/common"
)

type PgAddress struct {
	DB          *PgDb
	TableEntity TableEntity
}

func (addr PgAddress) Update(model *common.AddressModel) error {
	query := fmt.Sprintf("update %s set address_line1 = $1, address_line2 = $2, address_city = $3, address_postal_code = $4, address_district = $5 where id = $6", addr.TableEntity.Name)
	_, err := addr.DB.Database.Exec(query, model.Line1, model.Line2, model.City, model.PostalCode, model.District, addr.TableEntity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (addr PgAddress) Model() *common.AddressModel {
	return &common.AddressModel{}
}
