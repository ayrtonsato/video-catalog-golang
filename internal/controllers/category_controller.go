package controllers

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/helpers"
	"github.com/ayrtonsato/video-catalog-golang/internal/protocols"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
)

type GetCategoriesController struct {
	category services.ReaderCategory
}

func NewGetCategoriesController(category services.ReaderCategory) GetCategoriesController {
	return GetCategoriesController{
		category,
	}
}

func (c *GetCategoriesController) Handle() protocols.HttpResponse {
	listCategories, err := c.category.GetCategories()
	if err != nil {
		return helpers.HTTPInternalError()
	}
	return helpers.HTTPOk(listCategories)
}

type SaveCategoryController struct {
	category services.WriterCategory
	dto SaveCategoryDTO
	validation protocols.Validation
}

func NewSaveCategoryController(category services.WriterCategory,
	dto SaveCategoryDTO,
	validation protocols.Validation) SaveCategoryController {
	return SaveCategoryController{
		category:    category,
		dto: dto,
		validation: validation,
	}
}

func (c *SaveCategoryController) Handle() protocols.HttpResponse {
	err := c.validation.Validate()
	if err != nil {
		return helpers.HTTPBadRequestError(err)
	}
	category, err := c.category.Save(c.dto.Name, c.dto.Description)
	if err != nil {
		return helpers.HTTPInternalError()
	}
	return protocols.HttpResponse{
		Code: 201,
		Body: category,
	}
}
