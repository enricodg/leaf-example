package pkgConfig

import (
	"fmt"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewAppConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize application config: %+v", err)
	}
	if err := container.Provide(InitNewRelicConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize new relic config: %+v", err)
	}
	if err := container.Provide(NewSentryConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize sentry config: %+v", err)
	}
	return nil
}