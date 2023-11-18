package handler

import (
	"context"
	pkgKafka "kafka-go-implementation-consumer/pkg/kafka"
	"log"

	"github.com/segmentio/kafka-go"
)

type eventHandler struct {
	kafka *pkgKafka.Kafka
}

func InitEventHandler(kafka *pkgKafka.Kafka) {
	eventHandler := eventHandler{
		kafka: kafka,
	}

	eventHandler.TestTopic()
}

func (h *eventHandler) TestTopic() {
	messages := make(chan *kafka.Message)
	errors := make(chan error)
	ctx := context.Background()

	reader, err := h.kafka.ConsumeMessage(ctx, "test-group-consumer-id", "test-topic", "1", messages, errors)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			select {
			case m := <-messages:
				// TODO: use usecase
				log.Printf("Received message: %s\n", string(m.Value))
				if err := reader.CommitMessages(ctx, *m); err != nil {
					log.Printf("Failed to commit message: %v\n", err)
				}
			case err := <-errors:
				log.Printf("Received error: %v\n", err)
			}
		}
	}()
}
