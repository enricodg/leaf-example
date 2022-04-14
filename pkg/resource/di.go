package pkgResource

import (
	"fmt"
	"github.com/enricodg/leaf-example/pkg/resource/injection"
	"go.uber.org/dig"
)

func Register(container *dig.Container) error {
	if err := container.Provide(injection.NewLogger); err != nil {
		return fmt.Errorf("[DI] can not initialize logger: %+v", err)
	}
	if err := container.Provide(injection.NewMessageQueue); err != nil {
		return fmt.Errorf("[DI] can not initialize Messaging Queue: %+v", err)
	}
	if err := container.Provide(injection.NewWebClientFactory); err != nil {
		return fmt.Errorf("[DI] can not initialize Web Client Factory: %+v", err)
	}
	if err := container.Provide(injection.NewCache); err != nil {
		return fmt.Errorf("[DI] can not initialize Cache: %+v", err)
	}
	if err := container.Provide(injection.NewDatabaseSQL); err != nil {
		return fmt.Errorf("[DI] can not initialize SQL database: %+v", err)
	}
	if err := container.Provide(injection.InitTracer); err != nil {
		return fmt.Errorf("[DI] can not initialize tracer: %+v", err)
	}
	if err := container.Provide(injection.NewToken); err != nil {
		return fmt.Errorf("[DI] can not initialize token: %+v", err)
	}
	if err := container.Provide(injection.NewValidatorTranslator); err != nil {
		return fmt.Errorf("[DI] can not initialize validator translator: %+v", err)
	}
	if err := container.Provide(injection.NewValidator); err != nil {
		return fmt.Errorf("[DI] can not initialize validator: %+v", err)
	}
	return nil
}
