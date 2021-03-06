package health

import (
	"context"
	pkgResource "github.com/enricodg/leaf-example/pkg/resource"
	"net/http"
)

const (
	toggle = `toggle`
	status = `status`
)

func (c *Controller) Check(ctx context.Context) (int, map[string]map[string]interface{}) {
	var (
		httpStatus   = http.StatusOK
		dependencies = make(map[string]map[string]interface{})
	)
	dependencies = map[string]map[string]interface{}{
		pkgResource.Memcache: {toggle: c.Resource.ConfigMemcache.MemcacheEnable},
		pkgResource.Redis:    {toggle: c.Resource.ConfigRedis.RedisEnable},
		pkgResource.Kafka:    {toggle: c.Resource.ConfigKafka.KafkaEnable},
		pkgResource.MySQL:    {toggle: c.Resource.ConfigMySQL.MySqlEnable},
	}
	httpStatus, dependencies[pkgResource.Memcache] = ping(ctx, httpStatus, dependencies[pkgResource.Memcache], c.Resource.Cache.Memcache.Ping)
	httpStatus, dependencies[pkgResource.Redis] = ping(ctx, httpStatus, dependencies[pkgResource.Redis], c.Resource.Cache.Redis.Ping)
	httpStatus, dependencies[pkgResource.Kafka] = ping(ctx, httpStatus, dependencies[pkgResource.Kafka], c.Resource.MQ.Kafka.Ping)
	httpStatus, dependencies[pkgResource.MySQL] = ping(ctx, httpStatus, dependencies[pkgResource.MySQL], c.Resource.DatabaseSQL.MySQL.Ping)
	return httpStatus, dependencies
}

func ping(ctx context.Context, httpStatus int, data map[string]interface{}, fn func(ctx context.Context) error) (int, map[string]interface{}) {
	if activeValue, found := data[toggle]; found {
		if !activeValue.(bool) {
			data[status] = "DISABLED"
			return httpStatus, data
		}
	}

	if err := fn(ctx); err != nil {
		data[status] = "DOWN"
		return http.StatusInternalServerError, data
	}
	data[status] = "UP"
	return http.StatusOK, data
}
