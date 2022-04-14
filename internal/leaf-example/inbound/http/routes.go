package http

import (
	"github.com/enricodg/leaf-example/internal/leaf-example/inbound/http/health"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"github.com/labstack/echo/v4"
	leafHttpMiddleware "github.com/paulusrobin/leaf-utilities/appRunner/middleware/http"
	leafPrivilege "github.com/paulusrobin/leaf-utilities/common/constants/privilege"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Health   health.Controller
	Resource pkgResource.Resource
}

func (i Inbound) Routes(ec *echo.Echo, logger leafLogger.Logger) {
	ec.Use(leafHttpMiddleware.AppContextWithLogger(logger), leafHttpMiddleware.Tracer())
	ec.GET("/healthz/routes", i.Health.Routes).Name = leafPrivilege.Trusted
}