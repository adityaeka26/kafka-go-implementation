package kafka

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	kafkatrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/segmentio/kafka.go.v0"
)

type Kafka struct {
	writer   *kafkatrace.Writer
	sasl     bool
	brokers  []string
	username string
	password string
}

func NewKafka(sasl bool, hosts, username, password string, datadogEnable bool) (*Kafka, error) {
	// producer
	config := kafka.WriterConfig{}
	config.Brokers = strings.Split(hosts, ",")
	config.Balancer = &kafka.LeastBytes{}
	if sasl {
		mechanism, err := scram.Mechanism(scram.SHA512, username, password)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		config.Dialer = &kafka.Dialer{SASLMechanism: mechanism}
	}

	writer := kafkatrace.NewWriter(config, kafkatrace.WithServiceName("spbe-perizinan-event-kafka"))
	writer.AllowAutoTopicCreation = true

	return &Kafka{
		writer:   writer,
		sasl:     sasl,
		brokers:  strings.Split(hosts, ","),
		username: username,
		password: password,
	}, nil
}

func (k *Kafka) ConsumeMessage(ctx context.Context, groupId, topic, consumerId string) error {
	config := kafka.ReaderConfig{}
	config.Brokers = k.brokers
	config.GroupID = groupId
	config.Topic = topic
	config.MaxBytes = 10e6

	if k.sasl {
		mechanism, err := scram.Mechanism(scram.SHA512, k.username, k.password)
		if err != nil {
			return errors.WithStack(err)
		}

		config.Dialer = &kafka.Dialer{SASLMechanism: mechanism}
	}

	reader := kafkatrace.NewReader(config, kafkatrace.WithServiceName("spbe-perizinan-event-kafka"))

	var err error
	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("message at consumerId:%s topic:%v partition:%v offset:%v message:%s\n", consumerId, m.Topic, m.Partition, m.Offset, string(m.Value))
		if err := reader.CommitMessages(ctx, m); err != nil {
			log.Println(err)
			break
		}
	}

	return err
}

func (k *Kafka) Close(ctx context.Context) error {
	return k.writer.Close()
}
