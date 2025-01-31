package goPubSub

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"backend/internal/2_adapter/controller"
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
func NewKafkaConsumer() *kafka.Consumer {
	var consumer *kafka.Consumer
	var err error
	maxRetries := 20

	for i := 0; i < maxRetries; i++ {
		consumer, err = kafka.NewConsumer(
			&kafka.ConfigMap{
				"bootstrap.servers": "kafka:9092",
				// "bootstrap.servers": "localhost:9092",
				// "bootstrap.servers": "0.0.0.0:9092",
				"group.id":          "my-group",
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

// Start ....
func (receiver *GoPubSub) Start() {
	go receiver.subscribeOtherTopic()
	receiver.subscribeTestTopic()
}

// subscribeTestTopic ....
func (receiver *GoPubSub) subscribeTestTopic() {
	consumer := NewKafkaConsumer()
	err := consumer.Subscribe("test-topic", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("subscribeTestTopic Consumer started, waiting for messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("subscribeTestTopic Received message: %s\n", string(msg.Value))
			// RequestContextを生成してコントローラーに渡す
			// ctx := pubsubMiddleware.HeaderToContext(msg.Headers)

			// receiver.Controller.GetPersonList(ctx)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// subscribeOtherTopic ....
func (receiver *GoPubSub) subscribeOtherTopic() {
	consumer := NewKafkaConsumer()
	err := consumer.Subscribe("other-topic", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("subscribeOtherTopic Consumer started, waiting for messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("subscribeOtherTopic Received: %s\n", string(msg.Value))

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
