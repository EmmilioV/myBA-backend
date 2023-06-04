// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	gateway2 "go.mod/domain/employee/gateway"
	"go.mod/domain/employer/gateway"
	"go.mod/domain/employer/usecase"
	"go.mod/infrastructure/application"
	"go.mod/infrastructure/database"
	"go.mod/infrastructure/gateway/employee"
	"go.mod/infrastructure/gateway/employer"
	"go.mod/infrastructure/http/webserver"
)

// Injectors from wire.go:

func CreateApplication() *application.Application {
	webServer := webserver.NewWebServer()
	settings := application.LoadApplicationSettings()
	dbSettings := application.GetDBSettings(settings)
	dbConnection := database.NewConnection(dbSettings)
	idbProvider := employer.NewDBProvider(dbConnection)
	gateways := gateway.NewGateways(idbProvider)
	idbInserter := employee.NewDBInserter(dbConnection)
	idbDeleter := employee.NewDBDeleter(dbConnection)
	idbUpdater := employee.NewDBUpdater(dbConnection)
	gatewayGateways := gateway2.NewGateways(idbInserter, idbDeleter, idbUpdater)
	useCases := usecase.NewUseCases(gateways, gatewayGateways)
	applicationApplication := application.NewApplication(webServer, settings, useCases)
	return applicationApplication
}
