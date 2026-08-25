package goPubSub

import (
	"context"
	"fmt"

	pubsubMiddleware "backend/internal/1_framework/middleware/pubsub"
	requestContextMiddleware "backend/internal/1_framework/middleware/request_context"
	"backend/internal/logger"
)

// Start ....
func (receiver *GoPubSub) Start() error {
	ctx := context.Background()
	go receiver.subscribeOtherTopic(ctx)

	return receiver.subscribeTestTopic(ctx)
}

// subscribeTestTopic ....
func (receiver *GoPubSub) subscribeTestTopic(
	ctx context.Context,
) error {
	topicName := receiver.testTopic
	consumer, err := NewKafkaConsumer(
		ctx,
		receiver.bootstrapServers,
		receiver.consumerGroupID,
	)
	if err != nil {
		return err
	}

	err = consumer.Subscribe(topicName, nil)
	if err != nil {
		return fmt.Errorf("subscribe topic %s: %w", topicName, err)
	}

	logger.Logging(ctx, fmt.Sprintf("%s consumer started", topicName))
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			logger.Logging(ctx, fmt.Sprintf("%s received message: %s", topicName, string(msg.Value)))
			// RequestContextを生成してコントローラーに渡す
			messageCtx := pubsubMiddleware.HeaderToContext(msg.Headers)
			requestContext := requestContextMiddleware.GetRequestContext(messageCtx)
			logger.Logging(messageCtx, map[string]interface{}{
				"traceID":          requestContext.TraceID().GetValue(),
				"requestStartTime": requestContext.RequestStartTime().GetValue(),
			})

			// receiver.Controller.GetPersonList(ctx)
		} else {
			logger.Logging(ctx, fmt.Errorf("consume topic %s: %w", topicName, err))
		}
	}
}

// subscribeOtherTopic ....
func (receiver *GoPubSub) subscribeOtherTopic(
	ctx context.Context,
) {
	consumer, err := NewKafkaConsumer(
		ctx,
		receiver.bootstrapServers,
		receiver.consumerGroupID,
	)
	if err != nil {
		logger.Logging(ctx, err)
		return
	}
	topicName := receiver.otherTopic
	err = consumer.Subscribe(topicName, nil)
	if err != nil {
		logger.Logging(ctx, fmt.Errorf("subscribe topic %s: %w", topicName, err))
		return
	}

	logger.Logging(ctx, fmt.Sprintf("%s consumer started", topicName))
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			logger.Logging(ctx, fmt.Sprintf("%s received message: %s", topicName, string(msg.Value)))

		} else {
			logger.Logging(ctx, fmt.Errorf("consume topic %s: %w", topicName, err))
		}
	}
}
