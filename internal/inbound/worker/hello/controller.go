package hello

import (
	"github.com/enricodg/leaf-example/internal/usecases"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	UseCase  usecases.UseCase
	Resource pkgResource.Resource
}
