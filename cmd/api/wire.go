//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	appointmentGatewayDomain "go.mod/domain/appointment/gateway"
	customerGatewayDomain "go.mod/domain/customer/gateway"
	employeeGatewayDomain "go.mod/domain/employee/gateway"
	employerGatewayDomain "go.mod/domain/employer/gateway"
	serviceGatewayDomain "go.mod/domain/service/gateway"

	appointmentUseCases "go.mod/domain/appointment/usecase"
	customerUseCases "go.mod/domain/customer/usecase"
	employeeUseCases "go.mod/domain/employee/usecase"
	employerUseCases "go.mod/domain/employer/usecase"

	"go.mod/infrastructure/application"
	"go.mod/infrastructure/database"

	appointmentGatewayInfra "go.mod/infrastructure/gateway/appointment"
	customerGatewayInfra "go.mod/infrastructure/gateway/customer"
	employeeGatewayInfra "go.mod/infrastructure/gateway/employee"
	employerGatewayInfra "go.mod/infrastructure/gateway/employer"
	serviceGatewayInfra "go.mod/infrastructure/gateway/service"

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
		employeeGatewayInfra.NewDBProvider,
		employeeGatewayInfra.NewDBDeleter,
		employeeGatewayInfra.NewDBUpdater,
		employeeGatewayDomain.NewGateways,

		customerGatewayInfra.NewDBInserter,
		customerGatewayInfra.NewDBProvider,
		customerGatewayDomain.NewGateways,

		appointmentGatewayInfra.NewDBInserter,
		appointmentGatewayInfra.NewDBProvider,
		appointmentGatewayDomain.NewGateways,

		serviceGatewayInfra.NewDBInserter,
		serviceGatewayInfra.NewDBProvider,
		serviceGatewayDomain.NewGateways,

		employerUseCases.NewUseCases,
		employeeUseCases.NewUseCases,
		customerUseCases.NewUseCases,
		appointmentUseCases.NewUseCases,
	)

	return new(application.Application)
}
