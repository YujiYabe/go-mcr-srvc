package pubsub_publisher

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	gatewayExternal "backend/internal/2_adapter/gateway/external"
	"backend/internal/logger"
)

// PubsubPublisher ...
type (
	PubsubPublisher struct {
		Conn             *kafka.Producer
		bootstrapServers string
		testTopic        string
		flushTimeoutMS   int
		sampleUserName   string
	}
)

// NewToPubSub ...
func NewToPubSub(
	ctx context.Context,
	bootstrapServers string,
	testTopic string,
	flushTimeoutMS int,
	sampleUserName string,
) (
	toPubSub gatewayExternal.ToPubSub,
	err error,
) {

	pubsubPublisher := &PubsubPublisher{
		bootstrapServers: bootstrapServers,
		testTopic:        testTopic,
		flushTimeoutMS:   flushTimeoutMS,
		sampleUserName:   sampleUserName,
	}
	if false {
		conn, err := open(ctx, bootstrapServers, 30)
		if err != nil {
			return nil, err
		}

		pubsubPublisher.Conn = conn
	}

	return pubsubPublisher, nil
}

func open(
	ctx context.Context,
	bootstrapServers string,
	count uint,
) (
	producer *kafka.Producer,
	err error,
) {
	var lastErr error
	for attempt := uint(0); attempt <= count; attempt++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		conn, err := kafka.NewProducer(
			&kafka.ConfigMap{
				"bootstrap.servers": bootstrapServers,
			},
		)
		if err == nil {
			return conn, nil
		}

		lastErr = err
		logger.Logging(ctx, err)

		if attempt == count {
			break
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryBackoff(attempt)):
		}
	}

	return nil, fmt.Errorf("retry count over: %w", lastErr)
}

func retryBackoff(
	attempt uint,
) (
	duration time.Duration,
) {
	backoff := time.Duration(attempt+1) * time.Second
	if backoff > 5*time.Second {
		return 5 * time.Second
	}
	return backoff
}
