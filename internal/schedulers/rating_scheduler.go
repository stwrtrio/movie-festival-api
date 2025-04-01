package schedulers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
	"github.com/stwrtrio/movie-festival-api/pkg/kafka"
)

type RatingEventScheduler struct {
	Consumer   kafka.Consumer
	RatingRepo repositories.RatingRepository
}

func (res *RatingEventScheduler) ProcessRatingEvents() {
	log.Println("Processing rating events from Kafka...")

	for {
		msg, err := res.Consumer.ReadMessage()
		if err != nil {
			log.Printf("Error reading Kafka message: %v\n", err)
			continue
		}

		var ratingEvent models.Rating
		if err := json.Unmarshal(msg.Value, &ratingEvent); err != nil {
			log.Printf("Failed to unmarshal rating event: %v", err)
			continue
		}

		log.Printf("Processing rating event: %+v\n", ratingEvent)

		err = res.RatingRepo.UpdateMovieRating(context.Background(), ratingEvent.MovieID)
		if err != nil {
			log.Printf("Failed to update movie rating: %v", err)
		}
	}
}
