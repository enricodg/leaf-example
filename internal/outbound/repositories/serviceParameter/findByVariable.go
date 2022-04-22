package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) FindByVariable(ctx context.Context, orm leafSql.ORM, variable string) (model.ServiceParameter, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindByVariable)
	defer span.Finish()

	var data model.ServiceParameter
	if err := orm.WithContext(ctx).
		Table(model.ServiceParameter{}.TableName()).
		Where("variable = ?", variable).
		Take(ctx, &data).Error(); err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error on findByVariable: %v", tag, err),
			leafLogger.WithAttr("variable", variable)))
		return model.ServiceParameter{}, err
	}

	return data, nil
}
