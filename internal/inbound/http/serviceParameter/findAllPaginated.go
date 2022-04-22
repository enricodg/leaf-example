package serviceParameter

import (
	"github.com/enricodg/leaf-example/internal/inbound/model"
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	"net/http"
)

func (c *Controller) FindAllPaginated(eCtx echo.Context) error {
	request, err := model.NewServiceParameterPagingRequest(eCtx)
	if err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusBadRequest, err)
	}

	response, err := c.UseCase.ServiceParameter.FindAllPaginated(eCtx.Request().Context(), request.ToUCModel())
	if err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusInternalServerError, err)
	}
	return leafHttpResponse.NewJSONResponse(eCtx, http.StatusOK, response)
}
