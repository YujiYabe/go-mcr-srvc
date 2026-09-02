package pubsub_publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
)

type UserMessage struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}

// PublishTestTopic ...
func (receiver *PubsubPublisher) PublishTestTopic(
	ctx context.Context,
) (
	err error,
) {
	if receiver.Conn == nil {
		err = fmt.Errorf("pubsub producer is not initialized")
		return
	}

	message := UserMessage{
		ID:        1,
		Name:      receiver.sampleUserName,
		Timestamp: time.Now(),
	}

	jsonData, err := json.Marshal(message)
	if err != nil {
		err = fmt.Errorf("marshal pubsub message: %w", err)
		return
	}

	// Add headers to the message
	headers := pubsubMiddleware.ContextToHeader(ctx)

	err = receiver.Conn.Produce(
		&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &receiver.testTopic,
				Partition: kafka.PartitionAny,
			},
			Value:   jsonData,
			Headers: headers, // Add headers here
		},
		nil,
	)

	if err != nil {
		err = fmt.Errorf("produce pubsub message: %w", err)
		return
	}

	// メッセージ送信を確実にするため、完了を待つ
	receiver.Conn.Flush(receiver.flushTimeoutMS)

	err = nil

	return
}
