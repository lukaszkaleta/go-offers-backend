package postgres

import (
	"naborly/internal/api/common"
	"naborly/internal/api/user"
)

type PgUserSettings struct {
	DB *PgDb
	ID int
}

func NewPgUserSettings(db *PgDb, id int) user.UserSettings {
	return &PgUserSettings{db, id}
}

func (pgUserSetting *PgUserSettings) Model() *user.UserSettingsModel {
	return nil
}

func (pgUserSetting *PgUserSettings) Radar() common.Radar {
	return nil
}
