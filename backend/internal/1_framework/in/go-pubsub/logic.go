package goPubSub

import (
	"fmt"
	"log"

	// "backend/internal/2_adapter/controller"
	// pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
	// groupObject "backend/internal/4_domain/group_object"
	pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
	groupObject "backend/internal/4_domain/group_object"
)

// Start ....
func (receiver *GoPubSub) Start() {
	go receiver.subscribeOtherTopic()
	receiver.subscribeTestTopic()
}

// subscribeTestTopic ....
func (receiver *GoPubSub) subscribeTestTopic() {
	topicName := "test-topic"
	consumer := NewKafkaConsumer()

	err := consumer.Subscribe(topicName, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %s", err)
	}

	fmt.Println(topicName + " Consumer started, waiting for messages...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf(topicName+" Received message: %s\n", string(msg.Value))
			// RequestContextを生成してコントローラーに渡す
			ctx := pubsubMiddleware.HeaderToContext(msg.Headers)
			requestContext := groupObject.GetRequestContext(ctx)
			log.Println("== == == == == == == == == == ")
			log.Printf("%#v\n", requestContext.TraceID.GetValue())
			log.Printf("%#v\n", requestContext.RequestStartTime.GetValue())
			log.Println("== == == == == == == == == == ")

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
