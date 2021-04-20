package controllers

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/helpers"
	"github.com/ayrtonsato/video-catalog-golang/internal/protocols"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
	"github.com/gofrs/uuid"
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
	category   services.WriterCategory
	dto        SaveCategoryDTO
	validation protocols.Validation
}

func NewSaveCategoryController(category services.WriterCategory,
	dto SaveCategoryDTO,
	validation protocols.Validation) SaveCategoryController {
	return SaveCategoryController{
		category:   category,
		dto:        dto,
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
	return helpers.HTTPCreated(category)
}

type UpdateCategoryController struct {
	params     map[string]interface{}
	category   services.UpdateCategory
	dto        UpdateCategoryDTO
	validation protocols.Validation
}

func NewUpdateCategoryController(category services.UpdateCategory,
	dto UpdateCategoryDTO,
	validation protocols.Validation,
	params map[string]interface{}) UpdateCategoryController {
	return UpdateCategoryController{
		params:     params,
		category:   category,
		dto:        dto,
		validation: validation,
	}
}

func (u UpdateCategoryController) Handle() protocols.HttpResponse {
	newUUID := u.params["id"].(uuid.UUID)
	if newUUID == uuid.Nil {
		return helpers.HTTPNotFound()
	}
	err := u.validation.Validate()
	if err != nil {
		return helpers.HTTPBadRequestError(err)
	}
	err = u.category.Update(newUUID, u.dto.Name, u.dto.Description)
	if err != nil {
		if err == services.ErrNotFound {
			return helpers.HTTPNotFound()
		}
		return helpers.HTTPInternalError()
	}
	return helpers.HTTPOkNoContent()
}

type DeleteCategoryController struct {
	params   map[string]interface{}
	category services.DeleteCategory
}

func NewDeleteCategoryController(category services.DeleteCategory,
	params map[string]interface{}) DeleteCategoryController {
	return DeleteCategoryController{
		params:   params,
		category: category,
	}
}

func (u DeleteCategoryController) Handle() protocols.HttpResponse {
	newUUID := u.params["id"].(uuid.UUID)
	if newUUID == uuid.Nil {
		return helpers.HTTPNotFound()
	}
	err := u.category.Delete(newUUID)
	if err != nil {
		if err == services.ErrNotFound {
			return helpers.HTTPNotFound()
		}
		return helpers.HTTPInternalError()
	}
	return helpers.HTTPOkNoContent()
}

type GetSingleCategoryController struct {
	params   map[string]interface{}
	category services.ReaderCategory
}

func NewGetSingleCategoryController(category services.ReaderCategory,
	params map[string]interface{}) GetSingleCategoryController {
	return GetSingleCategoryController{
		params:   params,
		category: category,
	}
}

func (g GetSingleCategoryController) Handle() protocols.HttpResponse {
	newUUID := g.params["id"].(uuid.UUID)
	if newUUID == uuid.Nil {
		return helpers.HTTPNotFound()
	}
	category, err := g.category.GetCategory(newUUID)
	if err != nil {
		if err == services.ErrNotFound {
			return helpers.HTTPNotFound()
		}
		return helpers.HTTPInternalError()
	}
	return helpers.HTTPOk(category)
}
