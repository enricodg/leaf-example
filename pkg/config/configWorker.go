package pkgConfig

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"time"
)

type (
	WorkerConfig struct {
		// - Worker
		WorkerExampleInterval time.Duration `envconfig:"WORKER_EXAMPLE_INTERVAL" default:"180s"`
	}
)

func NewWorkerConfig() (WorkerConfig, error) {
	configuration := WorkerConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return WorkerConfig{}, err
	}
	return configuration, nil
}
