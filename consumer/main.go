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
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "1")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "1")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "2")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "3")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "4")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "5")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "6")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "7")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "8")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "9")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "1")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "1")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "2")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "3")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "4")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "5")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "6")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "7")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "8")
		if err != nil {
			log.Fatal(err)
		}
	}()

	go func() {
		err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "9")
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = kafka.ConsumeMessage(context.Background(), "test-group-1", "test", "10")
	if err != nil {
		log.Fatal(err)
	}
}
