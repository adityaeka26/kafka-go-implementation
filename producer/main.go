package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	pkgKafka "test-kraft/producer/pkg/kafka"
	"time"
)

func main() {
	kafka, err := pkgKafka.NewKafka(
		false,
		"localhost:9092,localhost:9093,localhost:9094",
		"",
		"",
		false,
	)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(kafka)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			err := kafka.SendMessage(context.Background(), "test-topic", []byte(fmt.Sprintf("mantap gan baru %d", i)))
			if err != nil {
				log.Println(err)
			}
		}()
	}

	time.Sleep(15 * time.Second)
}
