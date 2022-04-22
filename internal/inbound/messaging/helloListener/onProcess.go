package helloListener

import (
	"context"
	pkgError "github.com/enricodg/leaf-example/pkg/constant/error"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
)

func (c Controller) Handler(ctx context.Context, message leafMQ.Message) error {
	var err error
	// validate message data
	if message.Data == nil {
		err = pkgError.ErrKafkaMessageDataInvalid
		c.Log.Error(leafLogger.BuildMessage(ctx, err.Error()))
		return err
	}

	// unmarshal data
	data := map[string]string{}
	err = json.Unmarshal(message.Data, &data)
	if err != nil {
		err = pkgError.ErrKafkaUnmarshalDataInvalid
		c.Log.Error(leafLogger.BuildMessage(ctx, err.Error()))
		return err
	}

	return c.UseCases.Forwarder.ForwardMessage(ctx, data["text"])
}
