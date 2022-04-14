package pkgResource

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	"github.com/enricodg/leaf-example/pkg/resource/injection"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafToken "github.com/paulusrobin/leaf-utilities/token"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	leafValidator "github.com/paulusrobin/leaf-utilities/validator/validator"
	leafWebClient "github.com/paulusrobin/leaf-utilities/webClient/webClient"
	"go.uber.org/dig"
)

const (
	Memcache = `memcache`
	Redis    = `redis`
	MySQL    = `mysql`
	Kafka    = `kafka`
)

type (
	Resource struct {
		dig.In

		// Config for application
		ConfigApp               pkgConfig.AppConfig
		ConfigProfiler          pkgConfig.ProfilerConfig
		ConfigMySQL             pkgConfig.MySQLConfig
		ConfigRedis             pkgConfig.RedisConfig
		ConfigMemcache          pkgConfig.MemcacheConfig
		ConfigKafka             pkgConfig.KafkaConfig
		ConfigWorker            pkgConfig.WorkerConfig
		ConfigNewRelic          pkgConfig.NewRelicConfig
		ConfigWebServiceExample pkgConfig.WebServiceExampleConfig
		ConfigToken             pkgConfig.TokenConfig

		// Database
		DatabaseSQL injection.SQL

		// Log for logger
		Log leafLogger.Logger

		// MQ is messaging queue struct holder
		MQ injection.MessageQueue

		// Cache is cache struct holder
		Cache injection.Cache

		// WebClientFactory is web client factory to produce WebClient
		WebClientFactory leafWebClient.Factory

		// Tracer is utilities tracer
		Tracer leafTracer.Tracer

		// Token is utilities tracer
		Token leafToken.Token

		// Translator
		Translator injection.Translator

		// Validator
		Validator leafValidator.Validator
	}
)
