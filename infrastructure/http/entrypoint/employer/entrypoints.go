package employer

import (
	"go.mod/infrastructure/application"
)

func Entrypoints(app *application.Application) {
	employerGroup := app.WebServer.Group("v1/employer")

	employerGroup.POST("employee/register", registerEmployee(app.EmployerUsecases.HireEmployee))
	employerGroup.DELETE("employee/remove/:id", removeEmployee(app.EmployerUsecases.UnhireEmployee))
	employerGroup.PUT("employee/update", updateEmployee(app.EmployerUsecases.UpdateEmployeeInfo))

	employerGroup.POST("customer/register", registerCustomer(app.EmployerUsecases.RegisterCustomer))

	employerGroup.GET("appointment/all-by-me/:id", searchWithAppointmentsByID(app.EmployerUsecases.SearchWithAppointmentsByID))
}
