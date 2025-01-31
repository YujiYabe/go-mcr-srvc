package pubsub_publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
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
	topic := "test-topic"

	// 構造体のメッセージを作成
	message := UserMessage{
		ID:        1,
		Name:      "Alice",
		Timestamp: time.Now(),
	}

	// JSON にエンコード
	jsonData, err := json.Marshal(message)
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %s", err)
	}

	// Kafka にメッセージを送信
	err = receiver.Conn.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: jsonData,
	}, nil)

	if err != nil {
		log.Printf("Failed to produce message: %s", err)
	} else {
		fmt.Printf("Produced message: %s\n", string(jsonData))
	}

	// メッセージ送信を確実にするため、完了を待つ
	receiver.Conn.Flush(1000)
	return
}
