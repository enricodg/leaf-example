package pkgConfig

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"time"
)

type (
	RedisConfig struct {
		// - Toggle
		RedisEnable bool `envconfig:"REDIS_ENABLE" default:"false"`

		// - Redis
		RedisHost         []string      `envconfig:"REDIS_HOST" required:"true"`
		RedisPassword     string        `envconfig:"REDIS_PASSWORD" required:"true"`
		RedisDB           int           `envconfig:"REDIS_DB" default:"0"`
		RedisPoolSize     int           `envconfig:"REDIS_POOL_SIZE" default:"100"`
		RedisMinIdleConns int           `envconfig:"REDIS_MIN_IDLE_CONNS" default:"10"`
		RedisReadOnly     bool          `envconfig:"REDIS_READ_ONLY" default:"true"`
		RedisDialTimeout  time.Duration `envconfig:"REDIS_DIAL_TIMEOUT" default:"1s"`
		RedisPoolTimeout  time.Duration `envconfig:"REDIS_POOL_TIMEOUT" default:"10s"`
		RedisReadTimeout  time.Duration `envconfig:"REDIS_READ_TIMEOUT" default:"1s"`
		RedisWriteTimeout time.Duration `envconfig:"REDIS_WRITE_TIMEOUT" default:"3s"`
		RedisMaxConnAge   time.Duration `envconfig:"REDIS_MAX_CONN_AGE" default:"3m"`
	}
)

func NewRedisConfig() (RedisConfig, error) {
	configuration := RedisConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return RedisConfig{}, err
	}
	return configuration, nil
}
