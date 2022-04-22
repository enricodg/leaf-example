package serviceParameter

import (
	"context"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) FindByVariable(ctx context.Context, variable string) (model.ServiceParameterUpsertResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindByVariable)
	defer span.Finish()

	orm := u.resource.DatabaseSQL.MySQL
	result, err := u.Repositories.ServiceParameter.FindByVariable(ctx, orm, variable)
	if err != nil {
		return model.ServiceParameterUpsertResponse{}, err
	}

	return model.ServiceParameterUpsertResponse{
		ID:          result.BaseAuditable.ID,
		Variable:    result.Variable,
		Value:       result.Value,
		Description: result.Description,
		Version:     result.BaseAuditable.Version.Int64,
	}, nil
}
