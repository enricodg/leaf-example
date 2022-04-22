package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) FindByID(ctx context.Context, orm leafSql.ORM, id uint64) (model.ServiceParameter, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindByID)
	defer span.Finish()

	var data model.ServiceParameter
	if err := orm.WithContext(ctx).
		Table(model.ServiceParameter{}.TableName()).
		Where("id = ?", id).
		Take(ctx, &data).Error(); err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error on findByID: %v", tag, err),
			leafLogger.WithAttr("ID", id)))
		return model.ServiceParameter{}, err
	}

	return data, nil
}
