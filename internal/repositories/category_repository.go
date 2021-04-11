package repositories

import (
	"context"
	"database/sql"

	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	logger "github.com/ayrtonsato/video-catalog-golang/pkg/logger"
)

type Category interface {
	GetCategories() ([]models.Category, error)
	Save(name string, description string) (models.Category, error)
}

type CategoryRepository struct {
	db  *sql.DB
	log logger.Logger
}

func NewCategoryRepository(db *sql.DB, log logger.Logger) CategoryRepository {
	return CategoryRepository{
		db, log,
	}
}

func (c *CategoryRepository) saveIntoCategory(row RepoReader) (models.Category, error) {
	var category models.Category
	err := row.Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.IsActive,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt)
	if err != nil {
		c.log.Error(err.Error())
		return models.Category{}, err
	}
	return category, nil
}

func (c *CategoryRepository) GetCategories() ([]models.Category, error) {
	var categories []models.Category
	rows, err := c.db.QueryContext(
		context.Background(), "SELECT id, name, description, is_active, created_at, updated_at, deleted_at FROM categories",
	)
	if err != nil {
		c.log.Error(err.Error())
		return []models.Category{}, err
	}
	for rows.Next() {
		newCategory, err := c.saveIntoCategory(rows)
		if err != nil {
			return []models.Category{}, err
		}
		categories = append(categories, newCategory)
	}
	if err := rows.Err(); err != nil {
		c.log.Error(err.Error())
		return []models.Category{}, err
	}
	if len(categories) == 0 {
		return make([]models.Category, 0), nil
	}
	return categories, nil
}

func (c *CategoryRepository) Save(name string, description string) (models.Category, error) {
	insertStatement := `INSERT INTO categories(name, description)
		VALUES($1, $2)
		RETURNING id, name, description, is_active, created_at, updated_at, deleted_at
	`
	stmt, err := c.db.Prepare(insertStatement)
	if err != nil {
		c.log.Error(err.Error())
		return models.Category{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(name, description)
	return c.saveIntoCategory(row)
}
