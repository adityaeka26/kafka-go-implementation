package cmd

import (
	"fmt"
	"kafka-go-implementation-consumer/cmd/rest"
	"kafka-go-implementation-consumer/config"
	"kafka-go-implementation-consumer/internal/handler"
	pkgKafka "kafka-go-implementation-consumer/pkg/kafka"
	"log"
)

func Execute() {
	config, err := config.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	kafka, err := pkgKafka.NewKafka(
		config.KafkaSasl,
		config.KafkaHosts,
		config.KafkaUsername,
		config.KafkaPassword,
		false,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(kafka)

	handler.InitEventHandler(kafka)

	if err := rest.ServeREST(config); err != nil {
		log.Fatal(err)
	}
}
