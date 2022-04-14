package pkgConfig

import (
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	"time"
)

type (
	AppConfig struct {
		// - Server
		ServiceVersion       string        `envconfig:"SERVICE_VERSION" required:"true"`
		ServiceName          string        `envconfig:"SERVICE_NAME" required:"true"`
		ServerPort           int           `envconfig:"SERVER_PORT" default:"8090" required:"true"`
		GracefullyDuration   time.Duration `envconfig:"GRACEFULLY_DURATION" default:"10s"`
		HealthCheckAccessKey string        `envconfig:"HEALTH_CHECK_ACCESS_KEY"`

		// - Interface Setting
		HttpEnable      bool `envconfig:"INTERFACE_HTTP_ENABLE" default:"true"`
		MessagingEnable bool `envconfig:"INTERFACE_MESSAGING_ENABLE" default:"true"`
		WorkerEnable    bool `envconfig:"INTERFACE_WORKER_ENABLE" default:"true"`

		// - Log
		LogLevel     string `envconfig:"LOG_LEVEL" default:"INFO"`
		LogFilePath  string `envconfig:"LOG_FILE_NAME"`
		LogFormatter string `envconfig:"LOG_FORMATTER" default:"JSON"`
	}
)

func NewAppConfig() (AppConfig, error) {
	configuration := AppConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return AppConfig{}, err
	}
	return configuration, nil
}
