package pkgConfig

import (
	"fmt"
	leafConfig "github.com/paulusrobin/leaf-utilities/config"
	leafKafka "github.com/paulusrobin/leaf-utilities/messageQueue/integrations/kafka"
	"time"
)

type (
	KafkaConfig struct {
		// - Toggle
		KafkaEnable bool `envconfig:"KAFKA_ENABLE" default:"false"`

		// - Kafka
		KafkaHost             []string           `envconfig:"KAFKA_HOST" required:"true"`
		KafkaClientID         string             `envconfig:"KAFKA_CLIENT_ID" required:"true"`
		KafkaConsumerGroup    string             `envconfig:"KAFKA_CONSUMER_GROUP" required:"true"`
		KafkaVersion          string             `envconfig:"KAFKA_VERSION" required:"true" default:"2.6.0"`
		KafkaConsumerWorker   int                `envconfig:"KAFKA_CONSUMER_WORKER" default:"100"`
		KafkaConsumerRetryMax int                `envconfig:"KAFKA_CONSUMER_RETRY_MAX" default:"10"`
		KafkaStrategy         leafKafka.Strategy `envconfig:"KAFKA_STRATEGY" default:"BalanceStrategyRoundRobin"`

		// - Kafka Slack Notification
		KafkaSlackNotificationActive  bool          `envconfig:"KAFKA_SLACK_NOTIFICATION_ACTIVE" default:"true"`
		KafkaSlackNotificationHook    string        `envconfig:"KAFKA_SLACK_NOTIFICATION_HOOK" default:"https://hooks.slack.com/services/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"`
		KafkaSlackNotificationTimeout time.Duration `envconfig:"KAFKA_SLACK_NOTIFICATION_TIMEOUT" default:"10s"`
	}
)

func NewKafkaConfig() (KafkaConfig, error) {
	configuration := KafkaConfig{}
	if err := leafConfig.NewFromEnv(&configuration); err != nil {
		return KafkaConfig{}, err
	}
	if len(configuration.KafkaHost) == 0 {
		return KafkaConfig{}, fmt.Errorf("kafka host is required")
	}
	return configuration, nil
}
