package services

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/stwrtrio/movie-festival-api/config"
	"github.com/stwrtrio/movie-festival-api/internal/models"
	"github.com/stwrtrio/movie-festival-api/internal/repositories"
)

type MovieService interface {
	CreateMovie(ctx context.Context, movie *models.Movie) error
	UpdateMovie(ctx context.Context, movie *models.Movie) error
	GetMovies(ctx context.Context, pagination models.PaginationRequest, useCache bool) (*models.PaginationResponse, error)
}

type movieService struct {
	repo     repositories.MovieRepository
	redis    config.Client
	cacheTTL time.Duration
}

func NewMovieService(repo repositories.MovieRepository, redis config.Client, cacheTTL int) MovieService {
	return &movieService{
		repo:     repo,
		redis:    redis,
		cacheTTL: time.Duration(cacheTTL) * time.Minute}
}

func (s *movieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return s.repo.CreateMovie(ctx, movie)
}

func (s *movieService) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	return s.repo.UpdateMovie(ctx, movie)
}

func (s *movieService) GetMovies(ctx context.Context, pagination models.PaginationRequest, useCache bool) (*models.PaginationResponse, error) {
	cacheKey := generateCacheKey("get-movies", pagination.Page, pagination.PageSize)
	if useCache {
		result, err := s.getMoviesFromCache(ctx, cacheKey)
		if err == nil && result != nil {
			return result, nil
		}
	}

	movies, total, err := s.repo.GetMovies(ctx, pagination)
	if err != nil {
		return nil, err
	}

	totalPages := total / int64(pagination.PageSize)
	if total%int64(pagination.PageSize) != 0 {
		totalPages++
	}

	response := &models.PaginationResponse{
		Page:       pagination.Page,
		PageSize:   pagination.PageSize,
		TotalItems: int(total),
		TotalPages: int(totalPages),
		Data:       movies,
	}

	go func() {
		_ = s.setCache(context.Background(), cacheKey, response, s.cacheTTL)
	}()

	return response, nil
}

func generateCacheKey(prefix string, page, pageSize int) string {
	return prefix + ":" + strconv.Itoa(page) + ":" + strconv.Itoa(pageSize)
}

func (s *movieService) getMoviesFromCache(ctx context.Context, key string) (*models.PaginationResponse, error) {
	val, err := s.redis.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var response models.PaginationResponse
	if err := json.Unmarshal([]byte(val), &response); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *movieService) setCache(ctx context.Context, key string, value *models.PaginationResponse, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return s.redis.Set(ctx, key, data, ttl).Err()
}
