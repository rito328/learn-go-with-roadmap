package main

import (
	"fmt"
	"log"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-kafka/v3/pkg/kafka"
	"github.com/ThreeDotsLabs/watermill/message"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	// Kafka の Publisher を作成
	publisher, err := kafka.NewPublisher(kafka.PublisherConfig{
		Brokers:   []string{"localhost:9092"},
		Marshaler: kafka.DefaultMarshaler{},
	}, logger)
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Close()

	topic := "example.topic"

	// メッセージを 5 回送信
	for i := 1; i <= 5; i++ {
		msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("Message %d", i)))
		if err := publisher.Publish(topic, msg); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Published: Message %d\n", i)
	}
}
