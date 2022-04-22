package serviceParameter

import (
	"context"
	"fmt"
	"github.com/enricodg/leaf-example/internal/outbound/model"
	leafSql "github.com/paulusrobin/leaf-utilities/database/sql/sql"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (r *repository) FindAllPaginated(ctx context.Context, orm leafSql.ORM, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindAllPaginated)
	defer span.Finish()

	sortMap := map[string]string{
		"variable":  "variable",
		"createdAt": "created_at",
	}

	var data []model.ServiceParameter
	paginationResult, err := orm.WithContext(ctx).
		Model(&model.ServiceParameter{}).
		PaginateData(ctx, &data, leafSql.PaginateOptions{
			Page:     request.Page,
			Limit:    request.Limit,
			Sort:     request.Sort,
			FieldMap: sortMap,
		})
	if err != nil {
		r.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s error on findAllPaginated: %v", tag, err)))
		return model.ServiceParameterPagingResponse{}, err
	}

	return model.ServiceParameterPagingResponse{
		BasePagingResponse: paginationResult,
		ServiceParameters:  data,
	}, nil
}
