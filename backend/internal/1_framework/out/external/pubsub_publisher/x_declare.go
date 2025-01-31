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
		Conn *kafka.Producer
	}
)

// NewToPubSub ...
func NewToPubSub() gatewayExternal.ToPubSub {
	ctx := context.Background()
	conn, err := open(ctx, 30)
	if err != nil {
		logger.Logging(ctx, err)
		panic(err)
	}

	pubsubPublisher := new(PubsubPublisher)
	pubsubPublisher.Conn = conn
	return pubsubPublisher

}

func open(
	ctx context.Context,
	count uint,
) (*kafka.Producer, error) {
	conn, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
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
