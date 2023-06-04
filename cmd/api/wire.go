//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	employeeGatewayDomain "go.mod/domain/employee/gateway"
	employerGatewayDomain "go.mod/domain/employer/gateway"
	employerUseCases "go.mod/domain/employer/usecase"
	"go.mod/infrastructure/application"
	"go.mod/infrastructure/database"
	employeeGatewayInfra "go.mod/infrastructure/gateway/employee"
	employerGatewayInfra "go.mod/infrastructure/gateway/employer"

	"go.mod/infrastructure/http/webserver"
)

func CreateApplication() *application.Application {
	wire.Build(
		webserver.NewWebServer,

		application.LoadApplicationSettings,
		application.NewApplication,
		application.GetDBSettings,

		database.NewConnection,

		employerGatewayInfra.NewDBProvider,
		employerGatewayDomain.NewGateways,

		employeeGatewayInfra.NewDBInserter,
		employeeGatewayInfra.NewDBDeleter,
		employeeGatewayInfra.NewDBUpdater,
		employeeGatewayDomain.NewGateways,

		employerUseCases.NewUseCases,
	)

	return new(application.Application)
}
