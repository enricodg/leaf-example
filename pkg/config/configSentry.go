package pkgConfig

import (
	"fmt"
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
)

type (
	SentryConfig struct {
		// - Sentry
		SentryEnable          bool    `envconfig:"SENTRY_ENABLE" default:"false"`
		SentryEnvironment     string  `envconfig:"SENTRY_ENVIRONMENT" default:"staging"`
		SentryDSN             string  `envconfig:"SENTRY_DSN" default:"https://342bf042ba0a4d72b2d535ed9646535a@newsentry.tanihub.com/21"`
		SentryAppName         string  `envconfig:"SENTRY_APP_NAME" default:"leaf-exmaple"`
		SentrySampleRate      float64 `envconfig:"SENTRY_SAMPLE_RATE"`
		SentryTraceSampleRate float64 `envconfig:"SENTRY_TRACE_SAMPLE_RATE"`
	}
)

func NewSentryConfig() (SentryConfig, error) {
	configuration := SentryConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return SentryConfig{}, err
	}

	if configuration.SentryEnable {
		if configuration.SentryDSN == "" {
			return SentryConfig{}, fmt.Errorf("setry dsn is required")
		}
	}

	return configuration, nil
}
