package main

import (
	"context"
	"log"

	"go.mod/infrastructure/http/entrypoint"
)

func main() {
	application := CreateApplication()

	entrypoint.SetEntrypoints(application)

	ctx := context.Background()

	if err := application.Run(ctx); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
