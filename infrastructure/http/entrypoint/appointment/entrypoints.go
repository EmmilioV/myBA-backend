package appointment

import "go.mod/infrastructure/application"

func Entrypoints(app *application.Application) {
	appointmentGroup := app.WebServer.Group("v1/appointment")

	appointmentGroup.POST("/add-service", addService(app.AppointmentUseCases.AddService))
	appointmentGroup.GET("/with-services/:id", searchWithServicesByID(app.AppointmentUseCases.SearchWithServicesByID))
}
