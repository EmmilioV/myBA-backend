package customer

import "go.mod/infrastructure/application"

func Entrypoints(app *application.Application) {
	employerGroup := app.WebServer.Group("v1/customer")

	employerGroup.POST("/schedule-an-appointment", scheduleAnAppointment(app.CustomerUsecases.ScheduleAnAppointment))
}
