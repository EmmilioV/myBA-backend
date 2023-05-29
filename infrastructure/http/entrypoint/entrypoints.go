package entrypoint

import (
	"go.mod/infrastructure/application"
	"go.mod/infrastructure/http/entrypoint/employer"
)

func SetEntrypoints(app *application.Application) {
	employer.Entrypoints(app)
}
