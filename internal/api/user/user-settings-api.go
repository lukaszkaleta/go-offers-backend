package user

import (
	"naborly/internal/api/common"
)

// API

type UserSettings interface {
	Model() *UserSettingsModel
	Radar() common.Radar
}

// Model

type UserSettingsModel struct {
	Radar *common.RadarModel `json:"radar"`
}

// Builder

// Solid

type SolidUserSettings struct {
	Id           int
	model        *UserSettingsModel
	UserSettings UserSettings
}

func (u SolidUserSettings) Model() *UserSettingsModel {
	return u.model
}

func (u SolidUserSettings) Radar() common.Radar {
	if u.UserSettings != nil {
		return common.NewSolidRadar(
			u.Model().Radar,
			u.UserSettings.Radar(),
		)
	}
	return common.NewSolidRadar(u.Model().Radar, nil)
}
