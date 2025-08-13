package postgres

import (
	"fmt"
	"naborly/internal/api/common"
)

type PgPosition struct {
	DB          *PgDb
	TableEntity TableEntity
}

func (pos PgPosition) Update(model *common.PositionModel) error {
	query := fmt.Sprintf("update %s set position_latitude = $1, position_longitude = $2 where id = $3", pos.TableEntity.Name)
	_, err := pos.DB.Database.Exec(query, model.Lat, model.Lon, pos.TableEntity.Id)
	if err != nil {
		return err
	}
	return nil
}

func (pos PgPosition) Model() *common.PositionModel {
	return &common.PositionModel{}
}
