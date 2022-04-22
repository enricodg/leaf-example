package hello

import (
	"context"
	"fmt"
	"github.com/paulusrobin/leaf-utilities/encoding/json"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
)

func (m *mHello) PublishMessage(ctx context.Context, message string) error {
	data := map[string]interface{}{
		"text": message,
	}
	byteData, err := json.Marshal(data)
	if err != nil {
		m.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s fail publish message, error marshall message %v", tag, err),
			leafLogger.WithAttr("data", data)))
		return err
	}

	msg := leafMQ.Message{
		Data: byteData,
	}
	if err := m.resource.MQ.Kafka.Publish(ctx, "hello.message", msg); err != nil {
		m.resource.Log.Error(leafLogger.BuildMessage(ctx,
			fmt.Sprintf("%s fail publish message %v", tag, err),
			leafLogger.WithAttr("data", data)))
		return err
	}

	return nil
}
