package pkgConfig

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"time"
)

type (
	WebServiceExampleConfig struct {
		Host            string        `envconfig:"EXAMPLE_HOST" required:"true"`
		TimeoutDuration time.Duration `envconfig:"EXAMPLE_TIMEOUT_DURATION" default:"3s"`
		RetryCount      int           `envconfig:"EXAMPLE_RETRY_COUNT" default:"0"`
	}
)

func NewWebServiceExampleConfig() (WebServiceExampleConfig, error) {
	configuration := WebServiceExampleConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return WebServiceExampleConfig{}, err
	}
	return configuration, nil
}
