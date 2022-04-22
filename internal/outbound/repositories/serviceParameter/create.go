package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) Create(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) (model.ServiceParameter, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingCreate)
	defer span.Finish()

	if err := orm.WithContext(ctx).
		Table(model.ServiceParameter{}.TableName()).
		Create(ctx, &serviceParameter).Error(); err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error create: %v", tag, err),
			leafLogger.WithAttr("serviceParameter", serviceParameter)))
		return model.ServiceParameter{}, err
	}

	return serviceParameter, nil
}
