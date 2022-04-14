package pkgConfig

import (
	"fmt"

	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(NewAppConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize application config: %+v", err)
	}
	if err := container.Provide(NewKafkaConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize kafka config: %+v", err)
	}
	if err := container.Provide(NewWorkerConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize worker config: %+v", err)
	}
	if err := container.Provide(NewMySQLConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize my sql config: %+v", err)
	}
	if err := container.Provide(NewRedisConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize redis config: %+v", err)
	}
	if err := container.Provide(NewMemcacheConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize memcache config: %+v", err)
	}
	if err := container.Provide(NewProfilerConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize profiler config: %+v", err)
	}
	if err := container.Provide(InitNewRelicConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize new relic config: %+v", err)
	}
	if err := container.Provide(NewSentryConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize sentry config: %+v", err)
	}
	if err := container.Provide(NewWebServiceExampleConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize web service example config: %+v", err)
	}
	if err := container.Provide(NewTokenConfig); err != nil {
		return fmt.Errorf("[DI] can not initialize token config: %+v", err)
	}
	return nil
}
