package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	leafTypes "github.com/paulusrobin/leaf-utilities/common/types"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMandatory "github.com/paulusrobin/leaf-utilities/mandatory"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) Update(ctx context.Context, request model.ServiceParameterUpsertRequest) (model.ServiceParameterUpsertResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingUpdate)
	defer span.Finish()

	mandatory := leafMandatory.FromContext(ctx)
	orm := u.resource.DatabaseSQL.MySQL
	current, err := u.Outbound.Repositories.ServiceParameter.FindByVariable(ctx, orm, request.Variable)
	if err != nil {
		return model.ServiceParameterUpsertResponse{}, err
	}

	current.Value = request.Value
	current.Description = request.Description
	current.UpdatedBy = leafTypes.NewNullString(mandatory.User().Email())
	if err := u.Outbound.Repositories.ServiceParameter.Update(ctx, orm, current); err != nil {
		u.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error on update: %v", tag, err)))
		return model.ServiceParameterUpsertResponse{}, err
	}

	return u.FindByVariable(ctx, request.Variable)
}
