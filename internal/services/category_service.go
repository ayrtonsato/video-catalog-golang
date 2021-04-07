package services

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
)

type ReaderCategory interface {
	GetCategories() ([]models.Category, error)
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

type WriterCategory interface {
	Save(name string, description string) (models.Category, error)
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