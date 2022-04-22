package serviceParameter

import (
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	"net/http"
)

func (c *Controller) Delete(eCtx echo.Context) error {
	variable := eCtx.Param("variable")

	if err := c.UseCase.ServiceParameter.Delete(eCtx.Request().Context(), variable); err != nil {
		return leafHttpResponse.NewJSONResponse(eCtx, http.StatusInternalServerError, err)
	}
	return leafHttpResponse.NewJSONResponse(eCtx, http.StatusOK, map[string]interface{}{
		"message":  "OK",
		"varaible": variable,
	})
}
