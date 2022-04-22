package serviceParameter

import (
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	"net/http"
)

func (c *Controller) FindByVariable(eCtx echo.Context) error {
	variable := eCtx.Param("variable")

	response, err := c.UseCase.ServiceParameter.FindByVariable(eCtx.Request().Context(), variable)
	if err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusInternalServerError, err)
	}
	return leafHttpResponse.NewJSONResponse(eCtx, http.StatusOK, response)
}
