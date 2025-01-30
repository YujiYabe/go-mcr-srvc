package goPubSub

import (
	"backend/internal/2_adapter/controller"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// GoPubSub ...
type GoPubSub struct {
	Controller    controller.ToController
	KafkaConsumer *kafka.Consumer
}

// NewGoPubSub ...
func NewGoPubSub(
	controller controller.ToController,
) *GoPubSub {
	goPubSub := &GoPubSub{
		Controller:    controller,
		KafkaConsumer: NewKafkaConsumer(),
	}
	return goPubSub
}

// NewKafkaConsumer ...
func NewKafkaConsumer() *kafka.Consumer {
	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "my-group",
			"auto.offset.reset": "earliest",
		},
	)
	if err != nil {
		log.Fatalf("Failed to create consumer: %s", err)
	}
	// defer consumer.Close()

	return consumer
}

// Start ....
func (receiver *GoPubSub) Start() {
	receiver.subscribeTestTopic()
	receiver.subscribeOtherTopic()

}

// subscribeTestTopic ....
func (receiver *GoPubSub) subscribeTestTopic() {
	err := receiver.KafkaConsumer.Subscribe("test-topic", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("Consumer started, waiting for messages...")
	for {
		msg, err := receiver.KafkaConsumer.ReadMessage(-1)
		 receiver.KafkaConsumer.
		
		
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))
			receiver.Controller.GetPersonList(ctx)
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// subscribeOtherTopic ....
func (receiver *GoPubSub) subscribeOtherTopic() {
	err := receiver.KafkaConsumer.Subscribe("other-topic", nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println("Consumer started, waiting for messages...")
	for {
		msg, err := receiver.KafkaConsumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received: %s\n", string(msg.Value))

		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
