package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	logger "github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gofrs/uuid"
	"time"
)

type GenreWithCategories struct {
	models.Genre
	Categories []models.Category
}

type GenreDB interface {
	GetGenres() ([]models.Genre, error)
	GetByID(id uuid.UUID) (models.Genre, error)
	GetGenreByIDWithCategories(id uuid.UUID) (models.Genre, error)
	Save(name string, categories []uuid.UUID) (models.Genre, error)
	Update(id uuid.UUID, fields []string, values ...interface{}) error
	Delete(genre models.Genre) error
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

func (g *GenreRepository) saveIntoGenres(row RepoReader) (models.Genre, error) {
	var genre models.Genre
	err := row.Scan(
		&genre.ID,
		&genre.Name,
		&genre.IsActive,
		&genre.CreatedAt,
		&genre.UpdatedAt,
		&genre.DeletedAt)
	if err != nil {
		g.log.Error(err.Error())
		return models.Genre{}, err
	}
	return genre, nil
}

func (g *GenreRepository) GetGenres() ([]models.Genre, error) {
	var genres []models.Genre
	rows, err := g.db.QueryContext(
		context.Background(),
		"SELECT id, name, is_active, created_at, updated_at, deleted_at FROM genres WHERE is_active=false",
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			g.log.Error(err.Error())
			return []models.Genre{}, ErrNoResult
		}
		g.log.Error(err.Error())
		return []models.Genre{}, err
	}
	defer rows.Close()
	for rows.Next() {
		newGenre, err := g.saveIntoGenres(rows)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return []models.Genre{}, ErrNoResult
			}
			return []models.Genre{}, err
		}
		genres = append(genres, newGenre)
	}
	if err := rows.Err(); err != nil {
		g.log.Error(err.Error())
		return []models.Genre{}, err
	}
	if len(genres) == 0 {
		return make([]models.Genre, 0), nil
	}
	return genres, nil
}

func (g *GenreRepository) GetGenreByID(id uuid.UUID) (models.Genre, error) {
	query := "SELECT * FROM genres WHERE id=$1 AND is_active=false"
	row := g.db.QueryRow(query, id)
	genre, err := g.saveIntoGenres(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Genre{}, ErrNoResult
		}
		return models.Genre{}, err
	}
	return genre, nil
}

func (g *GenreRepository) Save(name string, categories []uuid.UUID) (models.Genre, error) {
	insertGenreStatement := `INSERT INTO genres(name)
		VALUES($1)
		RETURNING id, name, is_active, created_at, updated_at, deleted_at
	`
	insertRelationStatement := `INSERT INTO categories_genres(category_id, genre_id)
		VALUES($1, $2)
	`
	tx, err := g.db.BeginTx(context.Background(), nil)
	if err != nil {
		return models.Genre{}, ErrOnSave
	}
	stmt, err := tx.Prepare(insertGenreStatement)
	if err != nil {
		TransactionRollback(tx, g.log, err)
		return models.Genre{}, ErrOnSave
	}
	defer stmt.Close()
	row := stmt.QueryRow(name)
	genre, err := g.saveIntoGenres(row)
	if err != nil {
		TransactionRollback(tx, g.log, err)
		return genre, ErrOnSave
	}
	for _, category := range categories {
		stmtRelation, err := tx.Prepare(insertRelationStatement)
		if err != nil {
			g.log.Error(err.Error())
			TransactionRollback(tx, g.log, err)
			return models.Genre{}, ErrOnSave
		}
		defer stmtRelation.Close()
		_, err = stmtRelation.Exec(category, genre.ID)
		if err != nil {
			g.log.Error(err.Error())
			TransactionRollback(tx, g.log, err)
			return models.Genre{}, ErrOnSave
		}
	}
	if errCommit := TransactionCommit(tx, g.log); errCommit != nil {
		return models.Genre{}, ErrOnSave
	}
	return genre, nil
}

func (g *GenreRepository) Update(id uuid.UUID, fields []string, values ...interface{}) error {
	updateStmt, err := DynamicUpdateQuery("genres", fields)
	if err != nil {
		g.log.Error(err.Error())
	}
	stmt, err := g.db.Prepare(updateStmt)
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnUpdate
	}
	defer stmt.Close()
	values = append(values, id)
	exec, err := stmt.Exec(values...)
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnUpdate
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnUpdate
	}
	if affected > 0 {
		return nil
	}
	return ErrOnUpdate
}

func (g *GenreRepository) Delete(genre models.Genre) error {
	updateStmt, err := DynamicUpdateQuery("genres", []string{"is_active", "deleted_at"})
	tx, err := g.db.BeginTx(context.Background(), nil)
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnDelete
	}

	// soft delete genre first
	stmt, err := tx.Prepare(updateStmt)
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnUpdate
	}
	defer stmt.Close()
	_, err = stmt.Exec(false, time.Now().UTC(), genre.ID)
	if err != nil {
		g.log.Error(err.Error())
		return ErrOnDelete
	}

	// delete relationship btw category and genre
	query := `DELETE FROM categories_genres WHERE genre_id=$1`
	stmt, err = tx.Prepare(query)
	if err != nil {
		TransactionRollback(tx, g.log, err)
		return ErrOnDelete
	}
	defer stmt.Close()

	_, err = stmt.Exec(genre.ID)
	if err != nil {
		g.log.Error(err.Error())
		TransactionRollback(tx, g.log, err)
		return ErrOnDelete
	}
	if errCommit := TransactionCommit(tx, g.log); errCommit != nil {
		return ErrOnDelete
	}
	return nil
}
