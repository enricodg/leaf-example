package repositories

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/repositories/serviceParameter"
	"go.uber.org/dig"
)

type Repositories struct {
	dig.In

	ServiceParameter serviceParameter.Repository
}

func Register(container *dig.Container) error {
	if err := container.Provide(serviceParameter.New); err != nil {
		return fmt.Errorf("[DI] error provide service parameter repository: %+v", err)
	}
	return nil
}
