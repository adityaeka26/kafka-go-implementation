package kafka

import (
	"context"
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

func (k *Kafka) ConsumeMessage(ctx context.Context, groupId, topic string, messages chan<- *kafka.Message, errChan chan<- error) (*kafkatrace.Reader, error) {
	config := kafka.ReaderConfig{}
	config.Brokers = k.brokers
	config.GroupID = groupId
	config.Topic = topic

	if k.sasl {
		mechanism, err := scram.Mechanism(scram.SHA512, k.username, k.password)
		if err != nil {
			errChan <- err
			return nil, err
		}

		config.Dialer = &kafka.Dialer{SASLMechanism: mechanism}
	}

	reader := kafkatrace.NewReader(config, kafkatrace.WithServiceName("spbe-perizinan-event-kafka"))

	go func() {
		defer close(messages)
		defer close(errChan)
		for {
			m, err := reader.FetchMessage(ctx)
			// err = errors.New("asd")
			if err != nil {
				errChan <- err
				continue
			}
			messages <- &m
			// fmt.Printf("message at consumerId:%s topic:%s partition:%v offset:%v message:%s\n", consumerId, m.Topic, m.Partition, m.Offset, string(m.Value))
		}
	}()

	return reader, nil
}

func (k *Kafka) Close(ctx context.Context) error {
	return k.writer.Close()
}
