package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) Delete(ctx context.Context, orm leafSql.ORM, serviceParameter model.ServiceParameter) error {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingDelete)
	defer span.Finish()

	if err := orm.WithContext(ctx).
		Model(model.ServiceParameter{}).
		Set("deleted_by", serviceParameter.DeletedBy).
		Delete(ctx, &serviceParameter).Error(); err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error delete: %v", tag, err),
			leafLogger.WithAttr("ID", serviceParameter.BaseAuditable.ID)))
		return err
	}

	return nil
}
