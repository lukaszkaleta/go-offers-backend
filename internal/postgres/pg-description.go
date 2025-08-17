package postgres

import (
	"fmt"
	"naborly/internal/api/common"
)

type PgDescription struct {
	DB          *PgDb
	TableEntity TableEntity
}

func (p *PgDescription) Update(model *common.DescriptionModel) error {
	query := fmt.Sprintf("update %s set description_value = $1, description_image_url = $2 where id = $3", p.TableEntity.Name)
	_, err := p.DB.Database.Exec(query, model.Value, model.ImageUrl, p.TableEntity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PgDescription) Model() *common.DescriptionModel {
	return &common.DescriptionModel{}
}
