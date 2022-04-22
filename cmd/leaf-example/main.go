package main

import (
	"context"
	"github.com/enricodg/leaf-example/internal/inbound"
	pkgDi "github.com/enricodg/leaf-example/pkg/di"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	leafRunner "github.com/paulusrobin/leaf-utilities/appRunner"
	leafServer "github.com/paulusrobin/leaf-utilities/appRunner/server"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"go.uber.org/dig"
)

type (
	ContainerCall func() (*dig.Container, error)
	Invoke        func(container *dig.Container) error
	InvokeError   func(container *dig.Container, err error)
)

func Run(containerCall ContainerCall, invoke Invoke, onError InvokeError) error {
	container, err := containerCall()
	if err != nil {
		return err
	}
	if err := invoke(container); err != nil {
		onError(container, err)
	}
	return nil
}

func run(container *dig.Container) error {
	return container.Invoke(func(inbound inbound.Inbound, resource pkgResource.Resource) error {
		tracer.Start(resource.Tracer)

		leafRunner.New().
			With(leafServer.NewProfiler(resource.ConfigApp.ServiceName, resource.ConfigApp.ServiceVersion,
				leafServer.WithProfilerLogger(resource.Log),
				leafServer.WithProfilerEnable(resource.ConfigProfiler.GoogleProfilerEnable),
				leafServer.WithProfilerProjectID(resource.ConfigProfiler.GoogleProfilerProjectID),
			)).
			With(leafServer.NewHttp(resource.ConfigApp.ServiceName, resource.ConfigApp.ServiceVersion,
				leafServer.WithHttpPort(resource.ConfigApp.ServerPort),
				leafServer.WithHttpHealthAccessKey(resource.ConfigApp.HealthCheckAccessKey),
				leafServer.WithHttpLogger(resource.Log),
				leafServer.WithHttpHealthCheck(inbound.Http.Health.Check),
				leafServer.WithHttpRegister(inbound.Http.Routes),
				leafServer.WithHttpValidator(resource.Validator),
				leafServer.WithHttpEnable(resource.ConfigApp.HttpEnable),
			)).
			With(leafServer.NewMessaging(resource.MQ.Kafka,
				leafServer.WithConsumerLogger(resource.Log),
				leafServer.WithConsumerRegister(inbound.Messaging.Listener),
				leafServer.WithConsumerEnable(resource.ConfigApp.MessagingEnable),
			)).
			With(leafServer.NewWorker(
				leafServer.WithWorkerLogger(resource.Log),
				leafServer.WithWorkerRegister(inbound.Worker.Work),
				leafServer.WithWorkerEnable(resource.ConfigApp.WorkerEnable),
			)).
			Run()

		tracer.Stop()
		return nil
	})
}

func onError(container *dig.Container, err error) {
	_ = container.Invoke(func(logger leafLogger.Logger) error {
		if err != nil {
			logger.Info(leafLogger.BuildMessage(context.Background(), err.Error()))
		}
		return nil
	})
}

func main() {
	if err := Run(pkgDi.Container, run, onError); err != nil {
		panic(err)
	}
}
