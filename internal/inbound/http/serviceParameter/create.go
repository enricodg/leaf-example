package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/inbound/model"
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	"gorm.io/gorm"
	"net/http"
)

func (c *Controller) Create(eCtx echo.Context) error {
	var req model.ServiceParameterCreateRequest
	if err := eCtx.Bind(&req); err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	if err := eCtx.Validate(&req); err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	result, err := c.UseCase.ServiceParameter.Create(eCtx.Request().Context(), req.ToUCModel())
	if err != nil {
		status := http.StatusInternalServerError
		switch err {
		case gorm.ErrRecordNotFound:
			status = http.StatusNotFound
			break
		}
		return leafHttpResponse.NewJSONResponse(eCtx, status, err)
	}

	return leafHttpResponse.NewJSONResponse(
		eCtx,
		http.StatusOK,
		result,
	)
}
