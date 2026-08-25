package goPubSub

import (
	"log"
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
func NewKafkaConsumer() (
	consumer *kafka.Consumer,
) {
	consumer = &kafka.Consumer{}
	var err error
	maxRetries := 20

	for i := 0; i < maxRetries; i++ {
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
		time.Sleep(5 * time.Second)
	}
	if err != nil {
		log.Fatalf("Failed to create consumer after retries: %s", err)
	}

	return consumer
}
