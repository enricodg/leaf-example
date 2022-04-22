package usecases

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/usecases/serviceParameter"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	"go.uber.org/dig"
)

type UseCase struct {
	dig.In

	ServiceParameter serviceParameter.UseCase
}

func Register(container *dig.Container) error {
	if err := container.Provide(shared.New); err != nil {
		return fmt.Errorf("[DI] error provide shared use case: %+v", err)
	}
	if err := container.Provide(serviceParameter.New); err != nil {
		return fmt.Errorf("[DI] error provide service parameter use case: %+v", err)
	}
	return nil
}
