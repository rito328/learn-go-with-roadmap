package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	// Kafka の Subscriber を作成
	subscriber, err := kafka.NewSubscriber(kafka.SubscriberConfig{
		Brokers:       []string{"localhost:9092"},
		Unmarshaler:   kafka.DefaultMarshaler{},
		ConsumerGroup: "example-group",
	}, logger)
	if err != nil {
		log.Fatal(err)
	}
	defer subscriber.Close()

	topic := "example.topic"

	// Kafka からメッセージを購読
	messages, err := subscriber.Subscribe(context.Background(), topic)
	if err != nil {
		log.Fatal(err)
	}

	consumeMessages(messages)
}

// メッセージを受信して処理する
func consumeMessages(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("Received: %s\n", string(msg.Payload))
		msg.Ack()
	}
}
