package webservices

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/webservices/health"
	"go.uber.org/dig"
)

type WebServices struct {
	dig.In

	Health health.WebService
}

func Register(container *dig.Container) error {
	if err := container.Provide(health.New); err != nil {
		return fmt.Errorf("[DI] error provide health web service: %+v", err)
	}
	return nil
}
