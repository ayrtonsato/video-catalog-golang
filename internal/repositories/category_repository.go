package repositories

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	logger "github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gofrs/uuid"
)

type Category interface {
	GetCategories() ([]models.Category, error)
	Save(name string, description string) (models.Category, error)
	Update(id uuid.UUID, fields []string, values ...interface{}) error
	GetByID(id uuid.UUID) (models.Category, error)
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

func (c *CategoryRepository) Update(id uuid.UUID, fields []string, values ...interface{}) error {
	updateStmt, err := DynamicUpdateQuery("categories", fields)
	if err != nil {
		c.log.Error(err.Error())
	}
	stmt, err := c.db.Prepare(updateStmt)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}
	defer stmt.Close()
	values = append(values, id)
	exec, err := stmt.Exec(values...)
	if err != nil {
		c.log.Error(err.Error())
		return err
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		c.log.Error(err.Error())
		return err
	}
	if affected > 0 {
		return nil
	}
	return errors.New("repository: failed to update row")
}

func (c *CategoryRepository) GetByID(id uuid.UUID) (models.Category, error) {
	query := "SELECT * FROM categories WHERE id=$1"
	row := c.db.QueryRow(query, id)
	category, err := c.saveIntoCategory(row)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Category{}, ErrNoResult
		}
		return models.Category{}, err
	}
	return category, nil
}
