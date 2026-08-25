package pubsub_publisher

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	gatewayExternal "backend/internal/2_adapter/gateway/external"
	"backend/internal/env"
	"backend/internal/logger"
)

// PubsubPublisher ...
type (
	PubsubPublisher struct {
		Conn *kafka.Producer
	}
)

// NewToPubSub ...
func NewToPubSub(
	ctx context.Context,
) (
	gatewayExternal.ToPubSub,
	error,
) {

	pubsubPublisher := new(PubsubPublisher)
	if false {
		conn, err := open(ctx, 30)
		if err != nil {
			return nil, err
		}

		pubsubPublisher.Conn = conn
	}

	return pubsubPublisher, nil
}

func open(
	ctx context.Context,
	count uint,
) (*kafka.Producer, error) {
	conn, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": env.PubSubConfig.BootstrapServers,
		},
	)

	if err != nil {
		if count == 0 {
			logger.Logging(ctx, err)
			return nil, fmt.Errorf(
				"retry count over")
		}
		time.Sleep(time.Second)
		count--
		return open(ctx, count)
	}

	return conn, nil
}
