package application

import (
	"context"
	"fmt"
	"log"

	appointmentUsecase "go.mod/domain/appointment/usecase"
	customerUsecase "go.mod/domain/customer/usecase"
	employerUsecase "go.mod/domain/employer/usecase"
	"go.mod/infrastructure/http/webserver"
	"golang.org/x/sync/errgroup"
)

type Application struct {
	WebServer   *webserver.WebServer
	AppSettings *Settings

	EmployerUsecases    *employerUsecase.UseCases
	CustomerUsecases    *customerUsecase.UseCases
	AppointmentUseCases *appointmentUsecase.UseCases
}

func NewApplication(
	webServer *webserver.WebServer,
	appSettings *Settings,
	employerUsecases *employerUsecase.UseCases,
	customerUsecases *customerUsecase.UseCases,
	appointmentUseCases *appointmentUsecase.UseCases,
) *Application {
	return &Application{
		WebServer:           webServer,
		AppSettings:         appSettings,
		EmployerUsecases:    employerUsecases,
		CustomerUsecases:    customerUsecases,
		AppointmentUseCases: appointmentUseCases,
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
