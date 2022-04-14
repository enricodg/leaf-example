package health

import (
	"github.com/labstack/echo/v4"
	leafHttpResponse "github.com/paulusrobin/leaf-utilities/appRunner/response/http"
	leafModel "github.com/paulusrobin/leaf-utilities/common/model"
	"net/http"
	"sort"
)

func getListEndpoints(e *echo.Echo) leafModel.HealthEndpoints {
	var (
		routes    []*echo.Route             = e.Routes()
		routeLen  int                       = len(routes)
		endpoints leafModel.HealthEndpoints = make(leafModel.HealthEndpoints, routeLen)
	)

	for i := 0; i < routeLen; i++ {
		endpoints[i] = leafModel.HealthEndpoint{
			Method: routes[i].Method,
			Path:   routes[i].Path,
			Verify: routes[i].Name,
		}
	}

	sort.Slice(endpoints, func(i, j int) bool {
		if endpoints[i].Path < endpoints[j].Path {
			return true
		}

		if endpoints[i].Path > endpoints[j].Path {
			return false
		}

		return endpoints[i].Method < endpoints[j].Method
	})

	return endpoints
}

func (c *Controller) Routes(eCtx echo.Context) error {
	endpoints := getListEndpoints(eCtx.Echo())
	response := leafModel.HealthRoutesResponse{Routes: endpoints.String()}
	return leafHttpResponse.NewJSONResponse(eCtx, http.StatusOK, response)
}