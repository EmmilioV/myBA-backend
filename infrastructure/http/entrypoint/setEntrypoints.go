package entrypoint

import (
	"go.mod/infrastructure/application"
	"go.mod/infrastructure/http/entrypoint/customer"
	"go.mod/infrastructure/http/entrypoint/employer"
)

func SetEntrypoints(app *application.Application) {
	employer.Entrypoints(app)
	customer.Entrypoints(app)
}
