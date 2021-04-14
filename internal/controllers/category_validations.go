package controllers

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/helpers"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SaveCategoryValidation struct {
	dto *SaveCategoryDTO
}

func NewSaveCategoryValidation(dto *SaveCategoryDTO) SaveCategoryValidation {
	return SaveCategoryValidation{
		dto: dto,
	}
}

func (s SaveCategoryValidation) Validate() error {
	return validation.ValidateStruct(s.dto,
		validation.Field(&s.dto.Name, validation.Required, validation.Length(5, 254)),
		validation.Field(&s.dto.Description, validation.Required),
	)
}

type UpdateCategoryValidation struct {
	dto *UpdateCategoryDTO
}

func NewUpdateCategoryValidation(dto *UpdateCategoryDTO) UpdateCategoryValidation {
	return UpdateCategoryValidation{
		dto: dto,
	}
}

func (s UpdateCategoryValidation) Validate() error {
	return validation.ValidateStruct(s.dto,
		validation.Field(
			&s.dto.ID,
			validation.By(helpers.UUIDIsRequired)),
		validation.Field(
			&s.dto.Name,
			validation.Required.When(s.dto.Description == "").
				Error("cannot be blank when description is blank"),
		),
	)
}
