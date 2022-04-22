package model

import (
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	"github.com/labstack/echo/v4"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
)

type (
	ServiceParameterCreateRequest struct {
		Variable    string `json:"variable" validate:"required"`
		Value       string `json:"value" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	ServiceParameterUpdateRequest struct {
		Variable    string `json:"-"`
		Value       string `json:"value" validate:"required"`
		Description string `json:"description" validate:"required"`
	}

	ServiceParameterPagingRequest struct {
		leafModel.PagingParams
	}
)

func (r ServiceParameterCreateRequest) ToUCModel() ucModel.ServiceParameterUpsertRequest {
	return ucModel.ServiceParameterUpsertRequest{
		Variable:    r.Variable,
		Value:       r.Value,
		Description: r.Description,
	}
}

func (r ServiceParameterUpdateRequest) ToUCModel() ucModel.ServiceParameterUpsertRequest {
	return ucModel.ServiceParameterUpsertRequest{
		Variable:    r.Variable,
		Value:       r.Value,
		Description: r.Description,
	}
}

func NewServiceParameterPagingRequest(ctx echo.Context) (ServiceParameterPagingRequest, error) {
	var req ServiceParameterPagingRequest
	if err := ctx.Bind(&req); err != nil {
		return ServiceParameterPagingRequest{}, err
	}
	if err := ctx.Validate(&req); err != nil {
		return ServiceParameterPagingRequest{}, err
	}

	return ServiceParameterPagingRequest{
		PagingParams: leafModel.PagingParams{
			Page:  getIntegerParameterWithDefault(ctx, "page", 1),
			Limit: getIntegerParameterWithDefault(ctx, "limit", 20),
			Sort:  getSort(ctx, []string{"-createdAt"}),
		},
	}, nil
}

func (r ServiceParameterPagingRequest) ToUCModel() ucModel.ServiceParameterPagingRequest {
	return ucModel.ServiceParameterPagingRequest{
		PagingParams: leafModel.PagingParams{
			Page:   r.Page,
			Limit:  r.Limit,
			Sort:   r.Sort,
			Filter: r.Filter,
		},
	}
}
