package pubsub_publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
	"backend/internal/env"
)

type UserMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

// PublishTestTopic ...
func (receiver *PubsubPublisher) PublishTestTopic(
	ctx context.Context,
) {
	message := UserMessage{
		ID:        1,
		Name:      env.PubSubConfig.SampleUserName,
		Timestamp: time.Now(),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %s", err)
	}

	// Add headers to the message
	headers := pubsubMiddleware.ContextToHeader(ctx)

	err = receiver.Conn.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &env.PubSubConfig.TestTopic,
				Partition: kafka.PartitionAny,
			},
			Value:   jsonData,
			Headers: headers, // Add headers here
		},
		nil,
	)

	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	} else {
		fmt.Printf("Produced message: %s\n", string(jsonData))
	}

	// メッセージ送信を確実にするため、完了を待つ
	receiver.Conn.Flush(env.PubSubConfig.FlushTimeoutMS)
}
