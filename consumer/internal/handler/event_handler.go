package handler

import (
	"context"
	"log"
	pkgKafka "test-kraft/consumer/pkg/kafka"
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
	go func() {
		err := h.kafka.ConsumeMessage(context.Background(), "test-group-consumer-id", "test-topic", "1")
		if err != nil {
			log.Println(err)
		}
	}()
}
