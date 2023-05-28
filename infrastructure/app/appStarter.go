package app

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type AppStarter struct {
	WebServer *gin.Engine
}

func NewAppStarter(
	webServer *gin.Engine,
) *AppStarter {
	return &AppStarter{
		webServer,
	}
}

func (appStarter *AppStarter) Run(
	ctx context.Context,
) error {
	addr := fmt.Sprintf(":%d", 8080)

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if err := appStarter.WebServer.Run(addr); err != nil {
			return err
		}

		return nil
	})

	log.Printf("WebServer listening in port %s...\n", addr)

	return g.Wait()
}
