package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	logger "github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gofrs/uuid"
)

type GenreDB interface {
	GetCategories() ([]models.Genre, error)
	GetByID(id uuid.UUID) (models.Genre, error)
	Save(name string) (models.Genre, error)
	Update(id uuid.UUID, fields []string, values ...interface{}) error
}

type GenreRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewGenreRepository(db *sql.DB, log logger.Logger) GenreRepository {
	return GenreRepository{
		db, log,
	}
}

func (c *GenreRepository) saveIntoGenres(row RepoReader) (models.Genre, error) {
	var genre models.Genre
	err := row.Scan(
		&genre.Id,
		&genre.Name,
		&genre.IsActive,
		&genre.CreatedAt,
		&genre.UpdatedAt,
		&genre.DeletedAt)
	if err != nil {
		c.log.Error(err.Error())
		return models.Genre{}, err
	}
	return genre, nil
}

func (c *GenreRepository) GetGenres() ([]models.Genre, error) {
	var genres []models.Genre
	rows, err := c.db.QueryContext(
		context.Background(), "SELECT id, name, is_active, created_at, updated_at, deleted_at FROM genres",
	)
	if err != nil {
		c.log.Error(err.Error())
		return []models.Genre{}, err
	}
	for rows.Next() {
		newGenre, err := c.saveIntoGenres(rows)
		if err != nil {
			return []models.Genre{}, err
		}
		genres = append(genres, newGenre)
	}
	if err := rows.Err(); err != nil {
		c.log.Error(err.Error())
		return []models.Genre{}, err
	}
	if len(genres) == 0 {
		return make([]models.Genre, 0), nil
	}
	return genres, nil
}

func (c *GenreRepository) GetGenreByID(id uuid.UUID) (models.Genre, error) {
	query := "SELECT * FROM genres WHERE id=$1"
	row := c.db.QueryRow(query, id)
	genre, err := c.saveIntoGenres(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Genre{}, ErrNoResult
		}
		return models.Genre{}, err
	}
	return genre, nil
}
