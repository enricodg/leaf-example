package pkgDi

import (
	"github.com/enricodg/leaf-example/internal/leaf-example/outbound"
	"github.com/enricodg/leaf-example/internal/leaf-example/usecases"
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"go.uber.org/dig"
	"sync"
)

var (
	container *dig.Container
	once      sync.Once
)

func Container() (*dig.Container, error) {
	var outer error

	once.Do(func() {
		container = dig.New()

		if err := pkgConfig.Register(container); err != nil {
			outer = err
			return
		}

		if err := pkgResource.Register(container); err != nil {
			outer = err
			return
		}

		if err := outbound.Register(container); err != nil {
			outer = err
			return
		}

		if err := usecases.Register(container); err != nil {
			outer = err
			return
		}
	})

	if outer != nil {
		return nil, outer
	}

	return container, nil
}
