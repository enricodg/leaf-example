package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	"github.com/getsentry/sentry-go"
	leafNewRelicTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/newRelic"
	leafSentryTracer "github.com/paulusrobin/leaf-utilities/tracer/integrations/sentry"
	leafTracer "github.com/paulusrobin/leaf-utilities/tracer/tracer"
	"github.com/paulusrobin/leaf-utilities/tracer/tracer/tracer"
	"time"
)

func InitTracer(appConfig pkgConfig.AppConfig, nrConfig pkgConfig.NewRelicConfig, sentryConfig pkgConfig.SentryConfig) (leafTracer.Tracer, error) {
	if nrConfig.NewRelicEnable {
		return leafNewRelicTracer.InitTracing(
			leafNewRelicTracer.WithAppName(nrConfig.NewRelicAppName),
			leafNewRelicTracer.WithLicense(nrConfig.NewRelicLicense),
		)
	}

	if sentryConfig.SentryEnable {
		return leafSentryTracer.InitTracing(
			leafSentryTracer.WithSentryOptions(sentry.ClientOptions{
				Dsn:              sentryConfig.SentryDSN,
				ServerName:       sentryConfig.SentryAppName,
				Release:          appConfig.ServiceVersion,
				Environment:      sentryConfig.SentryEnvironment,
				SampleRate:       sentryConfig.SentrySampleRate,
				TracesSampleRate: sentryConfig.SentryTraceSampleRate,
			}),
			leafSentryTracer.WithDeferFlushDuration(12*time.Second),
		)
	}

	sentry.CaptureMessage("Capture Initialization")

	return tracer.NoopTracer(), nil
}
