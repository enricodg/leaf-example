package serviceParameter

import (
	"context"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) Delete(ctx context.Context, variable string) error {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingDelete)
	defer span.Finish()

	orm := u.resource.DatabaseSQL.MySQL
	serviceParam, err := u.Repositories.ServiceParameter.FindByVariable(ctx, orm, variable)
	if err != nil {
		return err
	}

	//mandatory := leafMandatory.FromContext(ctx)
	//serviceParam.DeletedBy = leafTypes.NewNullString(mandatory.User().Email())

	if err := u.Repositories.ServiceParameter.Delete(ctx, orm, serviceParam); err != nil {
		return err
	}

	return nil
}
