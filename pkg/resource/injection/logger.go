package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafLogrus "github.com/paulusrobin/leaf-utilities/logger/integrations/logrus"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
)

func NewLogger(config pkgConfig.AppConfig) (leafLogger.Logger, error) {
	formatter, err := leafLogrus.GetLoggerFormatter(config.LogFormatter)
	if err != nil {
		return nil, err
	}
	return leafLogrus.New(
		leafLogrus.WithLevel(leafLogger.GetLoggerLevel(config.LogLevel)),
		leafLogrus.WithFormatter(formatter),
		leafLogrus.WithLogFilePath(config.LogFilePath),
	)
}
