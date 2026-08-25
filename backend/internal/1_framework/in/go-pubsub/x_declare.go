package goPubSub

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"backend/internal/2_adapter/controller"
	"backend/internal/env"
	// pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
)

// GoPubSub ...
type GoPubSub struct {
	Controller controller.ToController
}

// NewGoPubSub ...
func NewGoPubSub(controller controller.ToController) *GoPubSub {
	return &GoPubSub{
		Controller: controller,
	}
}

// NewKafkaConsumer ...
// kafkaではtopic毎にシングルトンの為、Consumerインスタンスを共有できない
func NewKafkaConsumer(
	ctx context.Context,
) (
	consumer *kafka.Consumer,
	err error,
) {
	consumer = &kafka.Consumer{}
	maxRetries := 20

	for i := 0; i < maxRetries; i++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		consumer, err = kafka.NewConsumer(
			&kafka.ConfigMap{
				"bootstrap.servers": env.PubSubConfig.BootstrapServers,
				"group.id":          env.PubSubConfig.ConsumerGroupID,
				"auto.offset.reset": "earliest",
			},
		)
		if err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(5 * time.Second):
		}
	}
	if err != nil {
		return nil, fmt.Errorf("create kafka consumer after retries: %w", err)
	}

	return consumer, nil
}
