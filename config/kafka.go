package config

import (
	"log"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaConfig struct {
	Brokers string
	Topic   string
	GroupID string
}

type KafkaClient struct {
	Producer *kafka.Producer
	Consumer *kafka.Consumer
}

func InitKafka() (*KafkaClient, error) {
	// Load Kafka configuration
	config, err := LoadKafkaConfig()
	if err != nil {
		log.Fatalf("Error loading Kafka configuration: %v", err)
	}

	// Initialize producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
	})
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
		return nil, err
	}

	// Initialize consumer
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": config.Brokers,
		"group.id":          config.GroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		producer.Close()
		log.Fatalf("Error creating Kafka consumer: %v", err)
		return nil, err
	}

	log.Println("Kafka Producer & Consumer initialized")

	return &KafkaClient{Producer: producer, Consumer: consumer}, nil
}

func LoadKafkaConfig() (*KafkaConfig, error) {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "movie_rating_events"
	}

	groupID := os.Getenv("KAFKA_GROUP_ID")
	if groupID == "" {
		groupID = "rating_consumer_group"
	}

	return &KafkaConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	}, nil
}

func (k *KafkaClient) Close() {
	k.Producer.Close()
	k.Consumer.Close()
}
