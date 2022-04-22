package serviceParameter

import (
	"context"
	obModel "github.com/enricodg/leaf-example/internal/outbound/model"
	"github.com/enricodg/leaf-example/internal/usecases/model"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
)

func (u *ucServiceParameter) FindAllPaginated(ctx context.Context, request model.ServiceParameterPagingRequest) (model.ServiceParameterPagingResponse, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracingFindAllPaginated)
	defer span.Finish()

	orm := u.resource.DatabaseSQL.MySQL
	response, err := u.Repositories.ServiceParameter.FindAllPaginated(ctx, orm, obModel.ServiceParameterPagingRequest{
		PagingParams: leafModel.PagingParams{
			Page:   request.Page,
			Limit:  request.Limit,
			Sort:   request.Sort,
			Filter: request.Filter,
		},
	})
	if err != nil {
		return model.ServiceParameterPagingResponse{}, err
	}

	return model.ServiceParameterPagingResponse{
		BasePagingResponse: response.BasePagingResponse,
		ServiceParameters:  response.ServiceParameters,
	}, nil
}
