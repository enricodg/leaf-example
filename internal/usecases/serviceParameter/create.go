package serviceParameter

import (
	"context"
	"fmt"
	model2 "github.com/enricodg/leaf-example/internal/outbound/model"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) Create(ctx context.Context, request model.ServiceParameterUpsertRequest) (model.ServiceParameterUpsertResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingCreate)
	defer span.Finish()

	mandatory := leafMandatory.FromContext(ctx)
	orm := u.resource.DatabaseSQL.MySQL
	serviceParameter, err := u.Outbound.Repositories.ServiceParameter.Create(ctx, orm, model2.ServiceParameter{
		Variable:    request.Variable,
		Value:       request.Value,
		Description: request.Description,
		BaseAuditable: model2.BaseAuditable[uint64]{
			CreatedBy: mandatory.User().Email(),
		},
	})
	if err != nil {
		u.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error on create: %v", tag, err)))
		return model.ServiceParameterUpsertResponse{}, err
	}

	response := model.ServiceParameterUpsertResponse{
		ID:          serviceParameter.BaseAuditable.ID,
		Variable:    serviceParameter.Variable,
		Value:       serviceParameter.Value,
		Description: serviceParameter.Description,
		Version:     serviceParameter.BaseAuditable.Version.Int64,
	}
	if err := u.Outbound.Cache.ServiceParameter.Set(ctx, response); err != nil {
		u.resource.Log.Warn(leafLogger.BuildMessage(ctx,
			"error on set cache service parameter",
			leafLogger.WithAttr("variable", serviceParameter.Variable),
			leafLogger.WithAttr("err", err)))
	}

	return response, nil
}
