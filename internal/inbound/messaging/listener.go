package messaging

import (
	"github.com/enricodg/leaf-example/internal/inbound/messaging/helloListener"
	leafMessagingMiddleware "github.com/paulusrobin/leaf-utilities/appRunner/middleware/messaging"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
	"go.uber.org/dig"
)

type Inbound struct {
	dig.In

	HelloListener helloListener.Controller
}

func (i Inbound) Listener(consumer leafMQ.Consumer, logger leafLogger.Logger) {
	consumer.Use(leafMessagingMiddleware.AppContextWithLogger(logger), leafMessagingMiddleware.Tracer())

	if err := consumer.Subscribe("hello.world", i.HelloListener.Dispatcher()); err != nil {
		logger.StandardLogger().Errorf("[MESSAGING-SERVER] failed to subscribe messaging topic %s: %+v",
			"hello.world", err)
	}
}
