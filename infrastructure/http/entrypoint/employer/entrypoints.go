package employer

import (
	"go.mod/infrastructure/application"
)

func Entrypoints(app *application.Application) {
	employerGroup := app.WebServer.Group("v1/employer")

	employerGroup.POST("/register-employee", registerEmployee(app.EmployerUsecases.HireEmployee))
}
