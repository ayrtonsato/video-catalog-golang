package validators

import (
	"github.com/ayrtonsato/video-catalog-golang/internal/controllers"
	"github.com/go-ozzo/ozzo-validation/v4"
)

func SaveCategoryControllerValidator(dto *controllers.SaveCategoryDTO) error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Name, validation.Required, validation.Length(5, 254)),
		validation.Field(&dto.Description, validation.Required),
	)
}