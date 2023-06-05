package application

import (
	"context"
	"fmt"
	"log"

	appointmentUseCase "go.mod/domain/appointment/usecase"
	customerUseCase "go.mod/domain/customer/usecase"
	employeeUseCase "go.mod/domain/employee/usecase"
	employerUseCase "go.mod/domain/employer/usecase"
	"go.mod/infrastructure/http/webserver"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer   *webserver.WebServer
	AppSettings *Settings

	EmployerUsecases    *employerUseCase.UseCases
	CustomerUsecases    *customerUseCase.UseCases
	AppointmentUseCases *appointmentUseCase.UseCases
	EmployeeUseCases    *employeeUseCase.UseCases
}

func NewApplication(
	webServer *webserver.WebServer,
	appSettings *Settings,
	employerUsecases *employerUseCase.UseCases,
	customerUsecases *customerUseCase.UseCases,
	appointmentUseCases *appointmentUseCase.UseCases,
	employeeUseCases *employeeUseCase.UseCases,
) *Application {
	return &Application{
		WebServer:           webServer,
		AppSettings:         appSettings,
		EmployerUsecases:    employerUsecases,
		CustomerUsecases:    customerUsecases,
		AppointmentUseCases: appointmentUseCases,
		EmployeeUseCases:    employeeUseCases,
	}
}

func (application *Application) Run(
	ctx context.Context,
) error {
	addr := fmt.Sprintf(":%s", application.AppSettings.Port)

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := application.WebServer.Run(addr); err != nil {
			return err
		}

		return nil
	})

	log.Printf("WebServer listening in port %s...\n", addr)

	return g.Wait()
}
