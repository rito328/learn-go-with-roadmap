package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func main() {
	logger := watermill.NewStdLogger(false, false)

	// Go のチャンネルを使う Pub/Sub
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, logger)

	// 購読開始
	topic := "example.topic"
	messages, err := pubSub.Subscribe(context.Background(), topic)
	if err != nil {
		log.Fatal(err)
	}

	go processMessages(messages)

	// メッセージを送信
	publishMessage(pubSub, topic)

	// 少し待機して終了（本来はアプリが動作し続ける）
	time.Sleep(2 * time.Second)
}

// メッセージの処理（Subscriber）
func processMessages(messages <-chan *message.Message) {
	for msg := range messages {
		fmt.Printf("Received message: %s\n", string(msg.Payload))
		msg.Ack() // メッセージを確認
	}
}

// メッセージの送信（Publisher）
func publishMessage(pubSub message.Publisher, topic string) {
	msg := message.NewMessage(watermill.NewUUID(), []byte("Hello, Watermill!"))
	if err := pubSub.Publish(topic, msg); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Message published!")
}
