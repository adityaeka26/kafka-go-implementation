package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	pkgKafka "test-kraft/producer/pkg/kafka"
	"time"
)

func sendMessage(kafka *pkgKafka.Kafka, i int) error {
	err := kafka.SendMessageWithAutoTopicCreation(context.Background(), "test", []byte(fmt.Sprintf("mantap gan %d", i)))
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func main() {
	kafka, err := pkgKafka.NewKafka(
		false,
		"localhost:9092,localhost:9093,localhost:9094",
		"",
		"",
		false,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(kafka)

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			sendMessage(kafka, i)
		}()
	}
	time.Sleep(3 * time.Second)

	// err = kafka.SendMessageWithAutoTopicCreation(context.Background(), "test", []byte(fmt.Sprintf("mantap gan %d", 1)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
