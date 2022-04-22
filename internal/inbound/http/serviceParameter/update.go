package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/inbound/model"
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	"net/http"
)

func (c *Controller) Update(eCtx echo.Context) error {
	variable := eCtx.Param("variable")
	var req model.ServiceParameterUpdateRequest
	if err := eCtx.Bind(&req); err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	if err := eCtx.Validate(&req); err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	req.Variable = variable
	response, err := c.UseCase.ServiceParameter.Update(eCtx.Request().Context(), req.ToUCModel())
	if err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusInternalServerError, err)
	}
	return leafHttpResponse.NewJSONResponse(eCtx, http.StatusOK, response)
}
