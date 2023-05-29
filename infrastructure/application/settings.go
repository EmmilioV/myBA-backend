package application

import (
	"github.com/kelseyhightower/envconfig"
	"go.mod/infrastructure/database"
)

type Settings struct {
	Port       string `envconfig:"PORT"`
	DBSettings *database.DBSettings
}

func LoadApplicationSettings() *Settings {
	var settings Settings

	if err := envconfig.Process("", &settings); err != nil {
		panic(err)
	}

	return &settings
}

func GetDBSettings(settings *Settings) *database.DBSettings {
	return settings.DBSettings
}
