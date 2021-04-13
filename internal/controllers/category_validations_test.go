package controllers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSaveCategoryValidation_Validate(t *testing.T) {
	dto := SaveCategoryDTO{
		Name:        "test_valid",
		Description: "test_valid",
	}
	validator := NewSaveCategoryValidation(&dto)
	err := validator.Validate()
	require.NoError(t, err)

	dto2 := SaveCategoryDTO{
		Name:        "test_valid",
		Description: "",
	}
	validator = NewSaveCategoryValidation(&dto2)
	err = validator.Validate()
	require.Error(t, err)
	require.True(t, err.Error() == "description: cannot be blank.")

	dto3 := SaveCategoryDTO{
		Name:        "",
		Description: "test_valid",
	}
	validator = NewSaveCategoryValidation(&dto3)
	err = validator.Validate()
	require.Error(t, err)
	require.True(t, err.Error() == "name: cannot be blank.")

	dto4 := SaveCategoryDTO{
		Name:        "",
		Description: "",
	}
	validator = NewSaveCategoryValidation(&dto4)
	err = validator.Validate()
	require.Error(t, err)
	require.True(t, err.Error() == "description: cannot be blank; name: cannot be blank.")

	dto5 := SaveCategoryDTO{}
	validator = NewSaveCategoryValidation(&dto5)
	err = validator.Validate()
	require.Error(t, err)
	require.True(t, err.Error() == "description: cannot be blank; name: cannot be blank.")
}

func TestUpdateCategoryValidation_Validate(t *testing.T) {
	dto := UpdateCategoryDTO{
		Name: "teste",
	}
	validator := NewUpdateCategoryValidation(&dto)
	err := validator.Validate()
	require.NoError(t, err)

	dto2 := UpdateCategoryDTO{
		Description: "teste",
	}
	validator2 := NewUpdateCategoryValidation(&dto2)
	err2 := validator2.Validate()
	require.NoError(t, err2)

	dto3 := UpdateCategoryDTO{}
	validator3 := NewUpdateCategoryValidation(&dto3)
	err3 := validator3.Validate()
	require.Error(t, err3)
	require.True(t, err3.Error() == "name: cannot be blank when description is blank.")

	dto4 := UpdateCategoryDTO{
		Name:        "",
		Description: "",
	}
	validator4 := NewUpdateCategoryValidation(&dto4)
	err4 := validator4.Validate()
	require.Error(t, err4)
	require.True(t, err4.Error() == "name: cannot be blank when description is blank.")
}
