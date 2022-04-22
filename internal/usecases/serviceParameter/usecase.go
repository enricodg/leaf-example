package serviceParameter

import (
	"context"
	"github.com/enricodg/leaf-example/internal/outbound"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	"github.com/enricodg/leaf-example/internal/usecases/shared"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
)

const (
	tag = `[ServiceParameterUseCase]`

	tracingCreate           = "usecase.serviceParameter.Create"
	tracingFindAllPaginated = "usecase.serviceParameter.FindAllPaginated"
	tracingFindByVariable   = "usecase.serviceParameter.FindByVariable"
	tracingUpdate           = "usecase.serviceParameter.Update"
	tracingDelete           = "usecase.serviceParameter.Delete"
)

type (
	UseCase interface {
		Create(ctx context.Context, request model.ServiceParameterUpsertRequest) (model.ServiceParameterUpsertResponse, error)
		FindAllPaginated(ctx context.Context, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error)
		FindByVariable(ctx context.Context, variable string) (model.ServiceParameterUpsertResponse, error)
		Update(ctx context.Context, request model.ServiceParameterUpsertRequest) (model.ServiceParameterUpsertResponse, error)
		Delete(ctx context.Context, variable string) error
	}
	ucServiceParameter struct {
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
	return &ucServiceParameter{
		Outbound: outbound,
		shared:   shared,
		resource: resource,
	}
}
