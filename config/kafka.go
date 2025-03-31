package config

import (
	"os"
)

type KafkaConfig struct {
	Brokers string
	Topic   string
	GroupID string
}

func LoadKafkaConfig() *KafkaConfig {
	return &KafkaConfig{
		Brokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
		Topic:   getEnv("KAFKA_TOPIC", "movie_rating_events"),
		GroupID: getEnv("KAFKA_GROUP_ID", "rating_consumer_group"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
