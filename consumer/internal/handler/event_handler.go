package handler

import (
	"context"
	"encoding/json"
	"kafka-go-implementation-consumer/internal/dto"
	"kafka-go-implementation-consumer/internal/usecase"
	pkgKafka "kafka-go-implementation-consumer/pkg/kafka"
	"log"

	"github.com/segmentio/kafka-go"
)

type eventHandler struct {
	kafka           *pkgKafka.Kafka
	consumerUsecase usecase.ConsumerUsecase
}

func InitEventHandler(kafka *pkgKafka.Kafka, consumerUsecase usecase.ConsumerUsecase) {
	eventHandler := eventHandler{
		kafka:           kafka,
		consumerUsecase: consumerUsecase,
	}

	eventHandler.TestTopic()
}

func (h *eventHandler) TestTopic() {
	messages := make(chan *kafka.Message)
	errors := make(chan error)
	ctx := context.Background()

	reader, err := h.kafka.ConsumeMessage(ctx, "test-group-consumer-id", "test-topic-10", messages, errors)
	if err != nil {
		log.Println(err)
	}

	go func() {
		for {
			select {
			case m := <-messages:
				// log.Printf("Received message: %s\n", string(m.Value))
				var req dto.TestTopicReq
				err := json.Unmarshal(m.Value, &req)
				if err != nil {
					log.Printf("Failed to unmarshal message: %v\n", err)
				}

				log.Printf("partition: %v", m.Partition)
				err = h.consumerUsecase.TestTopic(ctx, req)
				if err != nil {
					log.Printf("Error in usecase: %v\n", err)
				}

				if err := reader.CommitMessages(ctx, *m); err != nil {
					log.Printf("Failed to commit message: %v\n", err)
				}
			case err := <-errors:
				log.Printf("Received error: %v\n", err)
			}
		}
	}()
}
