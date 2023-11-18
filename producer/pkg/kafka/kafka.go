package kafka

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	kafkatrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/segmentio/kafka.go.v0"
)

type Kafka struct {
	writer  *kafkatrace.Writer
	brokers []string
}

func NewKafka(sasl bool, hosts, username, password string, datadogEnable bool) (*Kafka, error) {
	// producer
	config := kafka.WriterConfig{}
	config.Brokers = strings.Split(hosts, ",")
	config.Balancer = &kafka.LeastBytes{}
	config.RequiredAcks = 1
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
		writer:  writer,
		brokers: strings.Split(hosts, ","),
	}, nil
}

func (k *Kafka) SendMessage(ctx context.Context, topic string, value []byte) error {
	fmt.Printf("send message to topic %s: %s\n", topic, string(value))
	return k.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: value,
	})
}

func (k *Kafka) SendMessageWithAutoTopicCreation(ctx context.Context, topic string, value []byte) error {
	var err error
	const retries = 3
	for i := 0; i < retries; i++ {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		err = k.writer.WriteMessages(ctx, kafka.Message{
			Topic: topic,
			Value: value,
		})
		if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
			time.Sleep(time.Millisecond * 250)
			continue
		}

		if err != nil {
			return err
		}
		break
	}

	fmt.Printf("send message (with topic creation) to topic %s: %s\n", topic, string(value))
	return nil
}

func (k *Kafka) ConsumeMessage(ctx context.Context, groupId, topic string) error {
	reader := kafkatrace.NewReader(kafka.ReaderConfig{
		Brokers:  k.brokers,
		GroupID:  groupId,
		Topic:    topic,
		MaxBytes: 10e6,
	})

	var err error
	for {
		m, err := reader.FetchMessage(ctx)
		if err != nil {
			break
		}
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		if err := reader.CommitMessages(ctx, m); err != nil {
			break
		}
	}

	return err
}

func (k *Kafka) Close(ctx context.Context) error {
	return k.writer.Close()
}
