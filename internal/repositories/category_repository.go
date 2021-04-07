package repositories

import (
	"context"
	"database/sql"

	"github.com/ayrtonsato/video-catalog-golang/internal/models"
)

type Category interface {
	GetCategories() ([]models.Category, error)
	Save(name string, description string) (models.Category, error)
}

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) CategoryRepository {
	return CategoryRepository{
		db,
	}
}

func (c *CategoryRepository) SetDB(db *sql.DB) {
	c.db = db
}

func (*CategoryRepository) saveIntoCategory(row RepoReader) (models.Category, error) {
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
		if err.Error() == sql.ErrNoRows.Error() {
			return models.Category{}, ErrNoRows
		}
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
		if err.Error() == sql.ErrNoRows.Error() {
			return []models.Category{}, ErrNoRows
		}
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
		return []models.Category{}, err
	}
	return categories, nil
}

/*func (c *CategoryRepository) GetCategoryByUUID(uuid uuid.UUID) (models.Category, error) {
	row := c.db.QueryRowContext(context.Background(), "SELECT * FROM categories WHERE id = ?", uuid)
	category := models.Category{}
	if err := row.Scan(&category.Id, &category.Name, &category.Description, &category.IsActive, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt); err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return models.Category{}, NoRows
		}
		return models.Category{}, nil
	}
	return category, nil
}*/

func (c *CategoryRepository) Save(name string, description string) (models.Category, error) {
	insertStatement := `INSERT INTO categories(name, description)
		VALUES($1, $2)
		RETURNING id, name, description, is_active, created_at, updated_at, deleted_at
	`
	stmt, err := c.db.Prepare(insertStatement)
	if err != nil {
		return models.Category{}, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(name, description)
	return c.saveIntoCategory(row)
}
