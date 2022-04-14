package injection

import (
	pkgConfig "github.com/enricodg/leaf-example/pkg/config"
	leafLogger "github.com/paulusrobin/leaf-utilities/logger/logger"
	leafKafka "github.com/paulusrobin/leaf-utilities/messageQueue/integrations/kafka"
	leafMQ "github.com/paulusrobin/leaf-utilities/messageQueue/messageQueue"
)

// MessageQueue is prepared for multi consumer
type (
	MessageQueue struct {
		Kafka leafMQ.MessageQueue
	}
	Publisher struct {
		Kafka leafMQ.Publisher
	}
)

func (m MessageQueue) Publisher() Publisher {
	return Publisher{
		Kafka: m.Kafka,
	}
}

func newKafka(config pkgConfig.KafkaConfig, logger leafLogger.Logger) (leafMQ.MessageQueue, error) {
	return leafKafka.New(
		leafKafka.WithHost(config.KafkaHost),
		leafKafka.WithClientID(config.KafkaClientID),
		leafKafka.WithConsumerGroup(config.KafkaConsumerGroup),
		leafKafka.WithKafkaVersion(config.KafkaVersion),
		leafKafka.WithConsumerWorker(config.KafkaConsumerWorker),
		leafKafka.WithConsumerRetryMax(config.KafkaConsumerRetryMax),
		leafKafka.WithStrategy(config.KafkaStrategy),
		leafKafka.WithLog(logger),
		leafKafka.WithSlackNotification(leafKafka.SlackNotification{
			Active:  config.KafkaSlackNotificationActive,
			Hook:    config.KafkaSlackNotificationHook,
			Timeout: config.KafkaSlackNotificationTimeout,
		}),
	)
}

func NewMessageQueue(kafkaConfig pkgConfig.KafkaConfig, logger leafLogger.Logger) (MessageQueue, error) {
	var (
		kafka leafMQ.MessageQueue = leafMQ.NoopQueue()
		err   error
	)

	if kafkaConfig.KafkaEnable {
		kafka, err = newKafka(kafkaConfig, logger)
		if err != nil {
			return MessageQueue{}, err
		}
	}

	return MessageQueue{
		Kafka: kafka,
	}, nil
}
