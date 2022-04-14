package health

import (
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	Resource pkgResource.Resource
}