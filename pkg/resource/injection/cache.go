package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafCache "github.com/paulusrobin/leaf-utilities/cache/cache"
	leafMemcache "github.com/paulusrobin/leaf-utilities/cache/integrations/memcache"
	leafRedis "github.com/paulusrobin/leaf-utilities/cache/integrations/redis"
)

// Cache is prepared for multi cache
type Cache struct {
	Redis    leafCache.Redis
	Memcache leafCache.Memcache
}

func NewCache(redisConfig pkgConfig.RedisConfig, memcacheConfig pkgConfig.MemcacheConfig) (Cache, error) {
	var (
		redis    leafCache.Redis    = leafCache.NoopRedis()
		memcache leafCache.Memcache = leafCache.NoopMemcache()
		err      error
	)

	if redisConfig.RedisEnable {
		redis, err = leafRedis.New(
			leafRedis.WithAddress(redisConfig.RedisHost),
			leafRedis.WithPassword(redisConfig.RedisPassword),
			leafRedis.WithDB(redisConfig.RedisDB),
			leafRedis.WithPoolSize(redisConfig.RedisPoolSize),
			leafRedis.WithReadOnly(redisConfig.RedisReadOnly),
			leafRedis.WithDialTimeout(redisConfig.RedisDialTimeout),
			leafRedis.WithPoolTimeout(redisConfig.RedisPoolTimeout),
			leafRedis.WithReadTimeout(redisConfig.RedisReadTimeout),
			leafRedis.WithWriteTimeout(redisConfig.RedisWriteTimeout),
			leafRedis.WithMaxConnAge(redisConfig.RedisMaxConnAge),
		)
		if err != nil {
			return Cache{}, err
		}
	}

	if memcacheConfig.MemcacheEnable {
		memcache = leafMemcache.New()
	}

	return Cache{
		Redis:    redis,
		Memcache: memcache,
	}, nil
}
