package postgres

import (
	"fmt"
	"naborly/internal/api/common"
)

type PgPrice struct {
	DB          *PgDb
	TableEntity TableEntity
}

func (p *PgPrice) Update(model *common.PriceModel) error {
	query := fmt.Sprintf("update %s set price_value = $1, price_currency = $2 where id = $3", p.TableEntity.Name)
	_, err := p.DB.Database.Exec(query, model.Value, model.Currency, p.TableEntity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PgPrice) Model() *common.PriceModel {
	return &common.PriceModel{}
}
