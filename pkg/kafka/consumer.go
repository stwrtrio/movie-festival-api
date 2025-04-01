package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Consumer interface {
	Subscribe(topics []string) error
	ReadMessage() (*kafka.Message, error)
	Close()
}

type kafkaConsumer struct {
	consumer *kafka.Consumer
}

func NewConsumer(brokers, groupID string) (Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &kafkaConsumer{consumer: c}, nil
}

func (kc *kafkaConsumer) Subscribe(topics []string) error {
	return kc.consumer.SubscribeTopics(topics, nil)
}

func (kc *kafkaConsumer) ReadMessage() (*kafka.Message, error) {
	msg, err := kc.consumer.ReadMessage(-1)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (kc *kafkaConsumer) Close() {
	kc.consumer.Close()
	log.Println("Kafka consumer closed")
}
