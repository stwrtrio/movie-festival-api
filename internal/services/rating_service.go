package services

import (
	"context"
	"encoding/json"
	"log"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"github.com/stwrtrio/movie-festival-api/pkg/kafka"
)

type RatingService interface {
	RateMovie(ctx context.Context, rating *models.Rating) error
}

type ratingService struct {
	ratingRepo    repositories.RatingRepository
	kafkaProducer kafka.Producer
	kafkaTopic    string
}

func NewRatingService(ratingRepo repositories.RatingRepository, producer kafka.Producer) RatingService {
	return &ratingService{ratingRepo: ratingRepo, kafkaProducer: producer}
}

func (s *ratingService) RateMovie(ctx context.Context, rating *models.Rating) error {
	err := s.ratingRepo.RateMovie(ctx, rating)
	if err != nil {
		return err
	}

	// Send rating event to Kafka
	event := &models.Event{
		ID:        rating.ID,
		UserID:    rating.UserID,
		MovieID:   rating.MovieID,
		Score:     rating.Score,
		Comment:   rating.Comment,
		CreatedAt: rating.CreatedAt.String(),
		UpdatedAt: rating.UpdatedAt.String(),
	}

	eventData, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to marshal event:", err)
		return err
	}

	if err := s.kafkaProducer.ProduceMessage(s.kafkaTopic, eventData); err != nil {
		log.Println("Failed to produce Kafka message:", err)
		return err
	}

	log.Println("Sent rating event to Kafka:", string(eventData))
	return nil
}
