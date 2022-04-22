package health

import (
	"context"
	"github.com/enricodg/leaf-example/internal/outbound"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

const (
	tag = `[Health]`
)

type (
	UseCase interface {
		CheckHealth(ctx context.Context) error
	}
	ucHealth struct {
		outbound.Outbound
		shared   shared.UseCase
		resource pkgResource.Resource
	}
)

func New(
	outbound outbound.Outbound,
	shared shared.UseCase,
	resource pkgResource.Resource,
) UseCase {
	return &ucHealth{
		Outbound: outbound,
		shared:   shared,
		resource: resource,
	}
}
