package pkgConfig

import (
	"fmt"
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
)

type (
	NewRelicConfig struct {
		// - NewRelic
		NewRelicEnable  bool   `envconfig:"NEW_RELIC_ENABLE" default:"false"`
		NewRelicAppName string `envconfig:"NEW_RELIC_APP_NAME"`
		NewRelicLicense string `envconfig:"NEW_RELIC_LICENSE"`
	}
)

func InitNewRelicConfig() (NewRelicConfig, error) {
	configuration := NewRelicConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return NewRelicConfig{}, err
	}

	if configuration.NewRelicEnable {
		if configuration.NewRelicAppName == "" || configuration.NewRelicLicense == "" {
			return NewRelicConfig{}, fmt.Errorf("newrelic app name / license is required")
		}
	}

	return configuration, nil
}