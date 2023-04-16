package producer

import (
	"context"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

type TollProducer struct {
	producer *kafka.Producer
}

func NewTollProducer(bootstrapServers string) (*TollProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		fmt.Printf("failed to create producer: %s\n", err)
		return nil, errors.Wrap(err, "failed creating producer")
	}
	return &TollProducer{
		producer: p,
	}, nil
}

func (p *TollProducer) Produce(ctx context.Context, topic string, payload []byte) error {

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(payload),
	}, nil)
	if err != nil {
		return errors.Wrap(err, "producing toll event")
	}
	return nil
}
