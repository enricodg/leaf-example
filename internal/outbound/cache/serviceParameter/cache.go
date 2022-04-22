package serviceParameter

import (
	"context"
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"github.com/enricodg/leaf-example/pkg/resource/injection"
)

type (
	Cache interface {
		Set(ctx context.Context, data ucModel.ServiceParameterUpsertResponse) error
		GetByVariable(ctx context.Context, variable string) (ucModel.ServiceParameterUpsertResponse, error)
		RemoveByVariable(ctx context.Context, variable string) error
	}
	cServiceParameter struct {
		cache    injection.Cache
		resource pkgResource.Resource
	}
)

func New(
	cache injection.Cache,
	resource pkgResource.Resource,
) Cache {
	return &cServiceParameter{
		cache:    cache,
		resource: resource,
	}
}
