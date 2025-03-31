package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Producer interface {
	ProduceMessage(topic string, message []byte) error
	Close()
}

type kafkaProducer struct {
	producer *kafka.Producer
}

func NewProducer(brokers string) (Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
	})
	if err != nil {
		return nil, err
	}

	return &kafkaProducer{producer: p}, nil
}

func (kp *kafkaProducer) ProduceMessage(topic string, message []byte) error {
	return kp.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: message,
	}, nil)
}

func (kp *kafkaProducer) Close() {
	kp.producer.Close()
}
