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

	if err := u.Outbound.Cache.ServiceParameter.RemoveByVariable(ctx, variable); err != nil {
		return err
	}

	if err := u.Repositories.ServiceParameter.Delete(ctx, orm, serviceParam); err != nil {
		return err
	}

	return nil
}
