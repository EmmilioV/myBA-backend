package employer

import (
	"go.mod/infrastructure/application"
)

func Entrypoints(app *application.Application) {
	employerGroup := app.WebServer.Group("v1/employer")

	employerGroup.POST("/register-employee", registerEmployee(app.EmployerUsecases.HireEmployee))
	employerGroup.DELETE("/remove-employee/:id", removeEmployee(app.EmployerUsecases.UnhireEmployee))
	employerGroup.PUT("/update-employee/:id", updateEmployee(app.EmployerUsecases.UpdateEmployeeInfo))

	employerGroup.POST("/register-customer", registerCustomer(app.EmployerUsecases.RegisterCustomer))
}
