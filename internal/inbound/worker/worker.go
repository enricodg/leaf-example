package worker

import (
	"github.com/enricodg/leaf-example/internal/inbound/worker/hello"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	leafWorkerMiddleware "github.com/paulusrobin/leaf-utilities/appRunner/middleware/worker"
	leafWorker "github.com/paulusrobin/leaf-utilities/appRunner/worker"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	Hello    hello.Controller
	Resource pkgResource.Resource
}

func (i Inbound) Work(runner *leafWorker.Runners, logger leafLogger.Logger) {
	runner.Add(leafWorker.NewRunner(
		"WORKER - Check Health",
		i.Resource.ConfigWorker.WorkerExampleInterval,
		i.Hello.Run,
		leafWorkerMiddleware.AppContextWithLogger(logger),
		leafWorkerMiddleware.Tracer(),
	))
}
