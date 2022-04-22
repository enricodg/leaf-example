package cache

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/cache/serviceParameter"
	"go.uber.org/dig"
)

type Cache struct {
	dig.In

	ServiceParameter serviceParameter.Cache
}

func Register(container *dig.Container) error {
	if err := container.Provide(serviceParameter.New); err != nil {
		return fmt.Errorf("[DI] error provide service parameter cache: %+v", err)
	}
	return nil
}
