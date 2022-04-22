package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) Update(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) error {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingCreate)
	defer span.Finish()

	if err := orm.WithContext(ctx).
		Model(&serviceParameter).
		Updates(ctx, model.ServiceParameter{
			Description:   serviceParameter.Description,
			Value:         serviceParameter.Value,
			BaseAuditable: model.BaseAuditable[uint64]{UpdatedBy: serviceParameter.UpdatedBy},
		}).Error(); err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error update: %v", tag, err),
			leafLogger.WithAttr("serviceParameter", serviceParameter)))
		return err
	}

	return nil
}
