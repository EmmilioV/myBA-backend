package employee

import "go.mod/infrastructure/application"

func Entrypoints(app *application.Application) {
	customerGroup := app.WebServer.Group("v1/employee")

	customerGroup.GET("/by-id/with-services/:id", getEmployeeWithServicesInfo(app.EmployeeUseCases.SearchByIDWithServices))
	customerGroup.PUT("/update-service", updateEmployeeService(app.EmployeeUseCases.UpdateServiceInfo))
}
