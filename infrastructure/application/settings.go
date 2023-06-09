package application

import (
	"github.com/kelseyhightower/envconfig"
	"go.mod/infrastructure/database"
	"go.mod/infrastructure/messaging"
)

type Settings struct {
	Port       string `envconfig:"PORT"`
	DBSettings *database.Settings
	MQSettings *messaging.Settings
}

func LoadApplicationSettings() *Settings {
	var settings Settings

	if err := envconfig.Process("", &settings); err != nil {
		panic(err)
	}

	return &settings
}

func GetDBSettings(settings *Settings) *database.Settings {
	return settings.DBSettings
}

func GetMQSettings(settings *Settings) *messaging.Settings {
	return settings.MQSettings
}
