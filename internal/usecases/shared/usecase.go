package shared

import (
	"github.com/enricodg/leaf-example/internal/outbound"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

type (
	UseCase interface {
	}
	ucShared struct {
		outbound.Outbound
		resource pkgResource.Resource
	}
)

func New(
	outbound outbound.Outbound,
	resource pkgResource.Resource,
) UseCase {
	return &ucShared{
		Outbound: outbound,
		resource: resource,
	}
}
