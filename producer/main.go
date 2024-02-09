package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

const (
	Topic              = "example"
	KafkaServerAddress = "kafka-svc.default.svc.cluster.local:9092"
	Message            = "ping"
)

func setupWriter() {
	w := &kafka.Writer{
		Addr:                   kafka.TCP(KafkaServerAddress),
		Topic:                  Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	fmt.Println("Started Kafka producer")

	for {
		err := w.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(time.Now().String()),
				Value: []byte(Message),
			},
		)
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := w.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
		fmt.Println("Successfully wrote message: ", Message)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	setupWriter()
}
