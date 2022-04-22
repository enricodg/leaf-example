package serviceParameter

import (
	"context"
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func (c *cServiceParameter) Set(ctx context.Context, data ucModel.ServiceParameterUpsertResponse) error {
	key := data.Variable

	if err := c.cache.Redis.Set(ctx, key, data); err != nil {
		c.resource.Log.Warn(leafLogger.BuildMessage(ctx, "error on set data to redis",
			leafLogger.WithAttr("err", err),
			leafLogger.WithAttr("data", data)))
		return err
	}
	return nil
}
