package pkgResource

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
	"go.uber.org/dig"
)

//const (
//	Memcache   = "memcache"
//	Redis      = "redis"
//	MySQL      = "mysql"
//	PostgreSQL = "postgresql"
//	Kafka      = "kafka"
//	MongoDB    = "mongodb"
//)

type (
	Resource struct {
		dig.In

		// Config for application
		ConfigApp               pkgConfig.AppConfig

		// Log for logger
		Log leafLogger.Logger

		// Tracer is utilities tracer
		Tracer leafTracer.Tracer

		// Validator
		Validator leafValidator.Validator
	}
)
