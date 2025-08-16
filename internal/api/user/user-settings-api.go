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

// Solid

type SolidUserSettings struct {
	Id           int
	model        *UserSettingsModel
	userSettings UserSettings
}

func NewSolidUserSettings(model *UserSettingsModel, userSettings UserSettings, id int) UserSettings {
	return &SolidUserSettings{
		Id:           id,
		model:        model,
		userSettings: userSettings,
	}
}

func (u SolidUserSettings) Model() *UserSettingsModel {
	return u.model
}

func (u SolidUserSettings) Radar() common.Radar {
	if u.userSettings != nil {
		return common.NewSolidRadar(
			u.Model().Radar,
			u.userSettings.Radar(),
		)
	}
	return common.NewSolidRadar(u.Model().Radar, nil)
}
