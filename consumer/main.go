package main

import (
	"context"
	"fmt"
	"log"
	pkgKafka "test-kraft/consumer/pkg/kafka"
)

func main() {
	kafka, err := pkgKafka.NewKafka(
		true,
		"localhost:9092,localhost:9093,localhost:9094",
		"username",
		"password",
		true,
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(kafka)

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group", "test", "a")
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = kafka.ConsumeMessage(context.Background(), "test-group", "test", "b")
	if err != nil {
		log.Fatal(err)
	}
}
