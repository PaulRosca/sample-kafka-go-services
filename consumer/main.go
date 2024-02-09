package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

const (
	ConsumerGroup      = "example-group"
	ConsumerTopic      = "example"
	KafkaServerAddress = "kafka-svc.default.svc.cluster.local:9092"
)

func setupConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{KafkaServerAddress},
		GroupID:  ConsumerGroup,
		Topic:    ConsumerTopic,
		MaxBytes: 10e6, // 10MB
	})
	fmt.Println("Started Kafka consumer")
	defer r.Close()
	ctx := context.Background()
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			break
		}

		fmt.Println("Fetched message", string(m.Key), string(m.Value))

		if err := r.CommitMessages(ctx, m); err != nil {
			log.Fatal("failed to commit messages:", err)
		}
	}
}

func main() {
	setupConsumer()
}
