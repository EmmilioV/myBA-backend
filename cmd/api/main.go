package main

import (
	"context"
	"log"
)

func main() {
	appStarter := CreateAppStarter()

	ctx := context.Background()

	if err := appStarter.Run(ctx); err != nil {
		log.Fatalf("FATAL: %s", err.Error())
	}
}
