package usecases

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/usecases/forwarder"
	"github.com/enricodg/leaf-example/internal/usecases/health"
	"github.com/enricodg/leaf-example/internal/usecases/serviceParameter"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	"go.uber.org/dig"
)

type UseCase struct {
	dig.In

	ServiceParameter serviceParameter.UseCase
	Health           health.UseCase
	Forwarder        forwarder.UseCase
}

func Register(container *dig.Container) error {
	if err := container.Provide(shared.New); err != nil {
		return fmt.Errorf("[DI] error provide shared use case: %+v", err)
	}
	if err := container.Provide(serviceParameter.New); err != nil {
		return fmt.Errorf("[DI] error provide service parameter use case: %+v", err)
	}
	if err := container.Provide(health.New); err != nil {
		return fmt.Errorf("[DI] error provide health use case: %+v", err)
	}
	if err := container.Provide(forwarder.New); err != nil {
		return fmt.Errorf("[DI] error provide forwarder use case: %+v", err)
	}
	return nil
}
