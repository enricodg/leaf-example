package http

import (
	"github.com/enricodg/leaf-example/internal/inbound/http/health"
	"github.com/enricodg/leaf-example/internal/inbound/http/serviceParameter"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"github.com/labstack/echo/v4"
	leafHttpMiddleware "github.com/paulusrobin/leaf-utilities/appRunner/middleware/http"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Health           health.Controller
	ServiceParameter serviceParameter.Controller
	Resource         pkgResource.Resource
}

func (i Inbound) Routes(ec *echo.Echo, logger leafLogger.Logger) {
	ec.Use(leafHttpMiddleware.AppContextWithLogger(logger), leafHttpMiddleware.Tracer())
	ec.GET("/healthz/routes", i.Health.Routes)

	v1 := ec.Group("/v1")
	sp := v1.Group("/service-parameters")
	sp.GET("", i.ServiceParameter.FindAllPaginated)
	sp.POST("", i.ServiceParameter.Create, leafHttpMiddleware.Token(i.Resource.Token))

	spVariable := sp.Group("/variable")
	spVariable.GET("/:variable", i.ServiceParameter.FindByVariable)
	spVariable.PATCH("/:variable", i.ServiceParameter.Update, leafHttpMiddleware.Token(i.Resource.Token))
	spVariable.DELETE("/:variable", i.ServiceParameter.Delete, leafHttpMiddleware.Token(i.Resource.Token))
}
