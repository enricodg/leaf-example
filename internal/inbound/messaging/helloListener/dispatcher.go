package helloListener

import (
	"github.com/enricodg/leaf-example/internal/usecases"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"go.uber.org/dig"
)

type Controller struct {
	dig.In

	UseCases usecases.UseCase
	Log      leafLogger.Logger
}

func (c Controller) Dispatcher() leafMQ.Dispatcher {
	dispatcher := leafMQ.NewSingleEventDispatcher()
	dispatcher.AddHandler(c.Handler, c.ErrorHandler)
	return dispatcher
}
