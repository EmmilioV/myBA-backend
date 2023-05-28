//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go.mod/infrastructure/app"
	"go.mod/infrastructure/http/webserver"
)

func CreateAppStarter() *app.AppStarter {
	wire.Build(
		webserver.NewWebServer,

		app.NewAppStarter,
	)

	return new(app.AppStarter)
}
