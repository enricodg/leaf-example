package serviceParameter

import (
	"context"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func (c *cServiceParameter) RemoveByVariable(ctx context.Context, variable string) error {
	if err := c.cache.Redis.Remove(ctx, variable); err != nil {
		c.resource.Log.Warn(leafLogger.BuildMessage(ctx, "error on remove service parameter data from redis",
			leafLogger.WithAttr("err", err),
			leafLogger.WithAttr("key", variable)))
		return err
	}

	return nil
}
