package goPubSub

import (
	"context"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"backend/internal/2_adapter/controller"
	// pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
)

// GoPubSub ...
type GoPubSub struct {
	Controller       controller.ToController
	bootstrapServers string
	consumerGroupID  string
	testTopic        string
	otherTopic       string
}

// NewGoPubSub ...
func NewGoPubSub(
	controller controller.ToController,
	bootstrapServers string,
	consumerGroupID string,
	testTopic string,
	otherTopic string,
) (
	goPubSub *GoPubSub,
) {
	return &GoPubSub{
		Controller:       controller,
		bootstrapServers: bootstrapServers,
		consumerGroupID:  consumerGroupID,
		testTopic:        testTopic,
		otherTopic:       otherTopic,
	}
}

// NewKafkaConsumer ...
// kafkaではtopic毎にシングルトンの為、Consumerインスタンスを共有できない
func NewKafkaConsumer(
	ctx context.Context,
	bootstrapServers string,
	consumerGroupID string,
) (
	consumer *kafka.Consumer,
	err error,
) {
	consumer = &kafka.Consumer{}
	maxRetries := 20

	for retryIndex := 0; retryIndex < maxRetries; retryIndex++ {
		if err := ctx.Err(); err != nil {
			return nil, err
		}

		consumer, err = kafka.NewConsumer(
			&kafka.ConfigMap{
				"bootstrap.servers": bootstrapServers,
				"group.id":          consumerGroupID,
				"auto.offset.reset": "earliest",
			},
		)
		if err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(retryBackoff(uint(retryIndex))):
		}
	}
	if err != nil {
		return nil, fmt.Errorf("create kafka consumer after retries: %w", err)
	}

	return consumer, nil
}

func retryBackoff(
	attempt uint,
) (
	duration time.Duration,
) {
	if attempt >= 4 {
		return 5 * time.Second
	}

	return time.Duration(attempt+1) * time.Second
}
