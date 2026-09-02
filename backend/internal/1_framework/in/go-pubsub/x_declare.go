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
	goPubSub = &GoPubSub{
		Controller:       controller,
		bootstrapServers: bootstrapServers,
		consumerGroupID:  consumerGroupID,
		testTopic:        testTopic,
		otherTopic:       otherTopic,
	}
	return
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
	err = nil
	consumer = &kafka.Consumer{}
	maxRetries := 20

	for retryIndex := 0; retryIndex < maxRetries; retryIndex++ {
		if returnedErr := ctx.Err(); returnedErr != nil {
			consumer, err = nil, returnedErr
			return //nolint:nakedret // Use the project-wide named return convention.
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
			consumer, err = nil, ctx.Err()
			return //nolint:nakedret // Use the project-wide named return convention.
		case <-time.After(retryBackoff(uint(retryIndex))):
		}
	}
	if err != nil {
		consumer, err = nil, fmt.Errorf("create kafka consumer after retries: %w", err)
		return //nolint:nakedret // Use the project-wide named return convention.
	}

	err = nil
	return //nolint:nakedret // Use the project-wide named return convention.
}

func retryBackoff(
	attempt uint,
) (
	duration time.Duration,
) {
	if attempt >= 4 {
		duration = 5 * time.Second
		return
	}

	duration = time.Duration(attempt+1) * time.Second
	return
}
