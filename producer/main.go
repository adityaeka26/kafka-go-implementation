package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	pkgKafka "test-kraft/producer/pkg/kafka"
	"time"
)

type TestTopicReq struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	kafka, err := pkgKafka.NewKafka(
		false,
		"localhost:9092",
		"",
		"",
		false,
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(kafka)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()

			message := TestTopicReq{
				Email:   "aditya.eka409@gmail.com",
				Message: fmt.Sprintf("Mantap gan %d", i),
			}
			messageByte, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
			}

			err = kafka.SendMessageWithAutoTopicCreation(context.Background(), "test-topic-10", messageByte)
			if err != nil {
				log.Println(err)
			}
		}()
	}

	time.Sleep(5 * time.Second)
}
