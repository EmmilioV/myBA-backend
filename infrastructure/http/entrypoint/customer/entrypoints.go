package customer

import "go.mod/infrastructure/application"

func Entrypoints(app *application.Application) {
	customerGroup := app.WebServer.Group("v1/customer")

	customerGroup.POST("/schedule-an-appointment", scheduleAnAppointment(app.CustomerUsecases.ScheduleAnAppointment))
}
