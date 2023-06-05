//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	appointmentGatewayDomain "go.mod/domain/appointment/gateway"
	customerGatewayDomain "go.mod/domain/customer/gateway"
	employeeGatewayDomain "go.mod/domain/employee/gateway"
	employerGatewayDomain "go.mod/domain/employer/gateway"

	customerUseCases "go.mod/domain/customer/usecase"
	employerUseCases "go.mod/domain/employer/usecase"

	"go.mod/infrastructure/application"
	"go.mod/infrastructure/database"

	appointmentGatewayInfra "go.mod/infrastructure/gateway/appointment"
	customerGatewayInfra "go.mod/infrastructure/gateway/customer"
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

		customerGatewayInfra.NewDBInserter,
		customerGatewayInfra.NewDBProvider,
		customerGatewayDomain.NewGateways,

		appointmentGatewayInfra.NewDBInserter,
		appointmentGatewayInfra.NewDBProvider,
		appointmentGatewayDomain.NewGateways,

		employerUseCases.NewUseCases,
		customerUseCases.NewUseCases,
	)

	return new(application.Application)
}
