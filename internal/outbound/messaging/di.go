package messaging

import (
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/messaging/hello"
	"go.uber.org/dig"
)

type Messaging struct {
	dig.In

	Hello hello.Messaging
}

func Register(container *dig.Container) error {
	if err := container.Provide(hello.New); err != nil {
		return fmt.Errorf("[DI] error provide hello messaging: %+v", err)
	}
	return nil
}
