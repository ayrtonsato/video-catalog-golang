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
	testCases := []struct {
		name string
		tc   func(t *testing.T)
	}{
		{
			name: "Return nil when name and description is valid",
			tc: func(t *testing.T) {
				dto := UpdateCategoryDTO{
					Name:        "teste",
					Description: "teste",
				}
				validator := NewUpdateCategoryValidation(&dto)
				err := validator.Validate()
				require.NoError(t, err)
			},
		},
		{
			name: "Return error when pass empty dto",
			tc: func(t *testing.T) {
				dto := UpdateCategoryDTO{}
				validator := NewUpdateCategoryValidation(&dto)
				err := validator.Validate()
				require.Error(t, err)
				require.True(t, err.Error() == "description: cannot be blank; name: cannot be blank.")
			},
		},
		{
			name: "Return error when name and description is empty",
			tc: func(t *testing.T) {
				dto4 := UpdateCategoryDTO{
					Name:        "",
					Description: "",
				}
				validator4 := NewUpdateCategoryValidation(&dto4)
				err4 := validator4.Validate()
				require.Error(t, err4)
				require.True(t, err4.Error() == "description: cannot be blank; name: cannot be blank.")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.tc(t)
		})
	}
}
