package pkgConfig

import leafConfig "github.com/paulusrobin/leaf-utilities/config"

type (
	MemcacheConfig struct {
		// - Toggle
		MemcacheEnable bool `envconfig:"MEMCACHE_ENABLE" default:"false"`
	}
)

func NewMemcacheConfig() (MemcacheConfig, error) {
	configuration := MemcacheConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return MemcacheConfig{}, err
	}
	return configuration, nil
}
