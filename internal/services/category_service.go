package services

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/gofrs/uuid"
	"time"
)

type ReaderCategory interface {
	GetCategories() ([]models.Category, error)
	GetCategory(id uuid.UUID) (models.Category, error)
}

type WriterCategory interface {
	Save(name string, description string) (models.Category, error)
}

type UpdateCategory interface {
	Update(id uuid.UUID, name string, description string) error
}

type DeleteCategory interface {
	Delete(id uuid.UUID) error
}

type GetCategoriesDbService struct {
	category repositories.Category
}

func NewGetCategoriesDbService(category repositories.Category) GetCategoriesDbService {
	return GetCategoriesDbService{
		category,
	}
}

func (g *GetCategoriesDbService) GetCategories() ([]models.Category, error) {
	return g.category.GetCategories()
}

func (g *GetCategoriesDbService) GetCategory(id uuid.UUID) (models.Category, error) {
	return g.category.GetByID(id)
}

type SaveDbCategoryService struct {
	category repositories.Category
}

func NewSaveDbCategoryService(category repositories.Category) SaveDbCategoryService {
	return SaveDbCategoryService{
		category,
	}
}

func (s *SaveDbCategoryService) Save(name string, description string) (models.Category, error) {
	return s.category.Save(name, description)
}

type UpdateDbCategoryService struct {
	category repositories.Category
}

func NewUpdateDbCategoryService(category repositories.Category) UpdateDbCategoryService {
	return UpdateDbCategoryService{
		category,
	}
}

func (n *UpdateDbCategoryService) Update(id uuid.UUID, name string, description string) error {
	_, err := n.category.GetByID(id)
	if err != nil {
		return ErrCategoryNotFound
	}
	err = n.category.Update(id, []string{"name", "description"}, name, description)
	if err != nil {
		return ErrCategoryUpdate
	}
	return nil
}

type DeleteDBCategoryService struct {
	category repositories.Category
}

func NewDeleteDBCategoryService(category repositories.Category) DeleteDBCategoryService {
	return DeleteDBCategoryService{
		category,
	}
}

func (d *DeleteDBCategoryService) Delete(id uuid.UUID) error {
	_, err := d.category.GetByID(id)
	if err != nil {
		return ErrCategoryNotFound
	}
	err = d.category.Update(id, []string{"is_active", "deleted_at"}, false, time.Now().UTC())
	if err != nil {
		return ErrCategoryUpdate
	}
	return nil
}
