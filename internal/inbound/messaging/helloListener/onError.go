package helloListener

import (
	"context"
	"fmt"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
)

func (c Controller) ErrorHandler(ctx context.Context, msg leafMQ.Message, err error) {
	c.Log.Error(leafLogger.BuildMessage(ctx, fmt.Sprintf(`[HelloListener] Error processing message, error: %+v`, err.Error())))
}
