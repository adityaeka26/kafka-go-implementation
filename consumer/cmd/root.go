package cmd

import (
	"fmt"
	"log"
	"test-kraft/consumer/cmd/rest"
	"test-kraft/consumer/config"
	"test-kraft/consumer/internal/handler"
	pkgKafka "test-kraft/consumer/pkg/kafka"
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
