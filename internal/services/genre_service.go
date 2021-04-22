package services

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/gofrs/uuid"
)

type ReaderGenre interface {
	GetGenres() ([]models.Category, error)
	GetGenreByID(id uuid.UUID) (models.Category, error)
}

type GetGenresDbService struct {
	genreRepository repositories.GenreRepository
}

func NewGetGenresDbService(genreRepository repositories.GenreRepository) GetGenresDbService {
	return GetGenresDbService{
		genreRepository,
	}
}

func (g *GetGenresDbService) GetGenres() ([]models.Genre, error) {
	return g.genreRepository.GetGenres()
}

func (g *GetGenresDbService) GetGenreByID(id uuid.UUID) (models.Genre, error) {
	genre, err := g.genreRepository.GetGenreByID(id)
	if err == repositories.ErrNoResult {
		return genre, ErrNotFound
	}
	return genre, err
}
