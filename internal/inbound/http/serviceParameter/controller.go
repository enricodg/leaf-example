package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/usecases"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	Resource pkgResource.Resource
	UseCase  usecases.UseCase
}
