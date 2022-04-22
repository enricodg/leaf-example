package pkgError

import "errors"

var (
	ErrOutbound                  = errors.New("error outbound")
	ErrKafkaMessageDataInvalid   = errors.New("kafka message data empty")
	ErrKafkaUnmarshalDataInvalid = errors.New("kafka message unmarshal failed")
)
