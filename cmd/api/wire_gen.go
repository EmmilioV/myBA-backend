// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	gateway4 "go.mod/domain/appointment/gateway"
	usecase3 "go.mod/domain/appointment/usecase"
	gateway3 "go.mod/domain/customer/gateway"
	usecase2 "go.mod/domain/customer/usecase"
	gateway2 "go.mod/domain/employee/gateway"
	usecase4 "go.mod/domain/employee/usecase"
	"go.mod/domain/employer/gateway"
	"go.mod/domain/employer/usecase"
	gateway5 "go.mod/domain/service/gateway"
	"go.mod/infrastructure/application"
	"go.mod/infrastructure/database"
	"go.mod/infrastructure/gateway/appointment"
	"go.mod/infrastructure/gateway/customer"
	"go.mod/infrastructure/gateway/employee"
	"go.mod/infrastructure/gateway/employer"
	"go.mod/infrastructure/gateway/service"
	"go.mod/infrastructure/http/webserver"
	"go.mod/infrastructure/messaging"
)

// Injectors from wire.go:

func CreateApplication() *application.Application {
	webServer := webserver.NewWebServer()
	settings := application.LoadApplicationSettings()
	databaseSettings := application.GetDBSettings(settings)
	connection := database.NewConnection(databaseSettings)
	idbProvider := employer.NewDBProvider(connection)
	gateways := gateway.NewGateways(idbProvider)
	idbInserter := employee.NewDBInserter(connection)
	gatewayIDBProvider := employee.NewDBProvider(connection)
	idbDeleter := employee.NewDBDeleter(connection)
	idbUpdater := employee.NewDBUpdater(connection)
	gatewayGateways := gateway2.NewGateways(idbInserter, gatewayIDBProvider, idbDeleter, idbUpdater)
	gatewayIDBInserter := customer.NewDBInserter(connection)
	idbProvider2 := customer.NewDBProvider(connection)
	gateways2 := gateway3.NewGateways(gatewayIDBInserter, idbProvider2)
	useCases := usecase.NewUseCases(gateways, gatewayGateways, gateways2)
	idbInserter2 := appointment.NewDBInserter(connection)
	idbProvider3 := appointment.NewDBProvider(connection)
	gateways3 := gateway4.NewGateways(idbInserter2, idbProvider3)
	usecaseUseCases := usecase2.NewUseCases(gateways, gateways2, gateways3)
	idbProvider4 := service.NewDBProvider(connection)
	idbInserter3 := service.NewDBInserter(connection)
	gatewayIDBUpdater := service.NewDBUpdater(connection)
	messagingSettings := application.GetMQSettings(settings)
	messagingConnection := messaging.NewConnection(messagingSettings)
	imqPublisher := service.NewMQPublisher(messagingConnection)
	gateways4 := gateway5.NewGateways(idbProvider4, idbInserter3, gatewayIDBUpdater, imqPublisher)
	useCases2 := usecase3.NewUseCases(gateways4, gateways, gateways3)
	useCases3 := usecase4.NewUseCases(gatewayGateways, gateways4)
	applicationApplication := application.NewApplication(webServer, settings, useCases, usecaseUseCases, useCases2, useCases3)
	return applicationApplication
}
