package services

import (
	"context"

	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	GetMovies(ctx context.Context, pagination models.PaginationRequest) (*models.PaginationResponse, error)
}

type movieService struct {
	repo repositories.MovieRepository
}

func NewMovieService(repo repositories.MovieRepository) MovieService {
	return &movieService{repo: repo}
}

func (s *movieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return s.repo.CreateMovie(ctx, movie)
}

func (s *movieService) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	return s.repo.UpdateMovie(ctx, movie)
}

func (s *movieService) GetMovies(ctx context.Context, pagination models.PaginationRequest) (*models.PaginationResponse, error) {
	movies, total, err := s.repo.GetMovies(ctx, pagination)
	if err != nil {
		return nil, err
	}

	totalPages := total / int64(pagination.PageSize)
	if total%int64(pagination.PageSize) != 0 {
		totalPages++
	}

	return &models.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalItems: int(total),
		TotalPages: int(totalPages),
		Data:       movies,
	}, nil
}
