package services

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/gofrs/uuid"
)

type ReaderGenre interface {
	GetGenres() ([]models.Category, error)
	GetGenreByID(id uuid.UUID) (models.Genre, error)
}

type GetGenresDBService struct {
	genreRepository repositories.GenreDB
}

func NewGetGenresDBService(genreRepository repositories.GenreDB) GetGenresDBService {
	return GetGenresDBService{
		genreRepository,
	}
}

func (g *GetGenresDBService) GetGenres() ([]models.Genre, error) {
	return g.genreRepository.GetGenres()
}

func (g *GetGenresDBService) GetGenreByID(id uuid.UUID) (models.Genre, error) {
	genre, err := g.genreRepository.GetByID(id)
	if err == repositories.ErrNoResult {
		return genre, ErrNotFound
	}
	return genre, err
}

type SaveGenre interface {
	Save(name string, categories []uuid.UUID) (models.Genre, error)
}

type SaveGenreDBService struct {
	genreRepository repositories.GenreDB
}

func NewSaveGenreDBService(genreRepository repositories.GenreDB) SaveGenreDBService {
	return SaveGenreDBService{
		genreRepository,
	}
}

func (s *SaveGenreDBService) Save(name string, categories []uuid.UUID) (models.Genre, error) {
	genre, err := s.genreRepository.Save(name, categories)
	if err != nil {
		return models.Genre{}, ErrSaveFailed
	}
	return genre, nil
}

type UpdateGenre interface {
	Update(id uuid.UUID, name string, categories []uuid.UUID) error
}

type UpdateGenreDBService struct {
	genreRepository repositories.GenreDB
}

func NewUpdateGenreDBService(genreRepository repositories.GenreDB) UpdateGenreDBService {
	return UpdateGenreDBService{
		genreRepository,
	}
}

func (u *UpdateGenreDBService) Update(id uuid.UUID, name string, categories []uuid.UUID) error {
	_, err := u.genreRepository.GetByID(id)
	if err != nil {
		return ErrNotFound
	}
	err = u.genreRepository.Update(id, categories, []string{"name"}, name)
	if err != nil {
		return ErrUpdateFailed
	}
	return nil
}

type DeleteGenre interface {
	Delete(id uuid.UUID) error
}

type DeleteGenreDBService struct {
	genreRepository repositories.GenreDB
}

func NewDeleteGenreDBService(genreRepository repositories.GenreDB) DeleteGenreDBService {
	return DeleteGenreDBService{
		genreRepository,
	}
}

func (d *DeleteGenreDBService) Delete(id uuid.UUID) error {
	err := d.genreRepository.
		Delete(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNoResult) {
			return ErrNotFound
		}
		return ErrUpdateFailed
	}
	return nil
}
