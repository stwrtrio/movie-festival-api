package services

import (
	"context"
	"errors"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
}

type movieService struct {
	repo repositories.MovieRepository
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return errors.New("not implemented")
}
