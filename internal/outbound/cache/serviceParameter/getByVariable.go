package serviceParameter

import (
	"context"
	ucModel "github.com/enricodg/leaf-example/internal/usecases/model"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func (c *cServiceParameter) GetByVariable(ctx context.Context, variable string) (ucModel.ServiceParameterUpsertResponse, error) {
	var response ucModel.ServiceParameterUpsertResponse
	if err := c.cache.Redis.Get(ctx, variable, &response); err != nil {
		c.resource.Log.Warn(leafLogger.BuildMessage(ctx, "error on get data from redis",
			leafLogger.WithAttr("err", err),
			leafLogger.WithAttr("variable", variable)))
		return ucModel.ServiceParameterUpsertResponse{}, err
	}
	return response, nil
}
