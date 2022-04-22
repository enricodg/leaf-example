package serviceParameter

import (
	"context"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	"github.com/go-redis/redis"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) FindByVariable(ctx context.Context, variable string) (model.ServiceParameterUpsertResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindByVariable)
	defer span.Finish()

	serviceParameterCached, err := u.Outbound.Cache.ServiceParameter.GetByVariable(ctx, variable)
	if err != redis.Nil && err != nil {
		return model.ServiceParameterUpsertResponse{}, err
	}

	if serviceParameterCached.ID != 0 {
		return serviceParameterCached, nil
	} else {
		orm := u.resource.DatabaseSQL.MySQL
		result, err := u.Repositories.ServiceParameter.FindByVariable(ctx, orm, variable)
		if err != nil {
			return model.ServiceParameterUpsertResponse{}, err
		}

		response := model.ServiceParameterUpsertResponse{
			ID:          result.BaseAuditable.ID,
			Variable:    result.Variable,
			Value:       result.Value,
			Description: result.Description,
			Version:     result.BaseAuditable.Version.Int64,
		}

		if err := u.Outbound.Cache.ServiceParameter.Set(ctx, response); err != nil {
			return model.ServiceParameterUpsertResponse{}, err
		}
		return response, nil
	}

}
