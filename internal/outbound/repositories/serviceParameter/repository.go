package serviceParameter

import (
	"context"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
)

const (
	tag = `[ServiceParameterRepository]`

	tracingFindAllPaginated = `outbound.repositories.serviceParameter.FindAllPaginated`
	tracingFindByVariable   = `outbound.repositories.serviceParameter.FindByVariable`
	tracingCreate           = `outbound.repositories.serviceParameter.Create`
	tracingUpdate           = `outbound.repositories.serviceParameter.Update`
	tracingDelete           = `outbound.repositories.serviceParameter.Delete`
	tracingFindByID         = `outbound.repositories.serviceParameter.FindByID`
)

type (
	Repository interface {
		FindAllPaginated(ctx context.Context, orm leafSql.ORM, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error)
		FindByVariable(ctx context.Context, orm leafSql.ORM, variable string) (model.ServiceParameter, error)
		FindByID(ctx context.Context, orm leafSql.ORM, id uint64) (model.ServiceParameter, error)
		Create(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) (model.ServiceParameter, error)
		Update(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) error
		Delete(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) error
	}

	repository struct {
		resource pkgResource.Resource
	}
)

func New(resource pkgResource.Resource) Repository {
	return &repository{
		resource: resource,
	}
}
