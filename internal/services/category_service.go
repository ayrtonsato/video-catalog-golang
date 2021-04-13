package services

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/gofrs/uuid"
)

type ReaderCategory interface {
	GetCategories() ([]models.Category, error)
}

type WriterCategory interface {
	Save(name string, description string) (models.Category, error)
}

type UpdateCategory interface {
	Update(id uuid.UUID, fields ...interface{}) error
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

func (n *UpdateDbCategoryService) Update(id uuid.UUID, fields []string, values ...interface{}) error {
	err := n.category.Update(id, fields, values...)
	if err != nil {
		return errors.New("service: failed to update category")
	}
	return nil
}
