package services

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	mock_repositories "github.com/ayrtonsato/video-catalog-golang/internal/repositories/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetCategoriesDbService_GetCategories(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	var fakeCategory = models.Category{
		Id: uid,
		Name: "valid_name",
		Description: "valid_description",
		IsActive: true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should call GetCategories once",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1)
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				_, _ = SUT.GetCategories()
			},
		},
		{
			name: "Should return an slice of models category without errors",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				listCategories := []models.Category{fakeCategory}
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1).
					Return(listCategories, nil)
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				result, err := SUT.GetCategories()
				require.NoError(t, err)
				require.Equal(t, result, listCategories)
			},
		},
		{
			name: "Should throw error if exists",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetCategories().
					Times(1).
					Return([]models.Category{}, errors.New("fake_error"))
				SUT := GetCategoriesDbService{
					category: ctgRepository,
				}
				_, err := SUT.GetCategories()
				require.NotEmpty(t, err)
				require.Equal(t, err.Error(), "fake_error")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tc.testCase(t, ctrl)
		})
	}
}

func TestSaveDbCategory_Save(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	var fakeCategory = models.Category{
		Id: uid,
		Name: "valid_name",
		Description: "valid_description",
		IsActive: true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should call Save Category repository correctly",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					Save(gomock.Eq("valid_name"), gomock.Eq("valid_description")).
					Times(1)
				SUT := SaveDbCategoryService{
					category: ctgRepository,
				}
				_, _ = SUT.Save("valid_name", "valid_description")
			},
		},
		{
			name: "Should return a model category without errors",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					Save(gomock.Eq("valid_name"), gomock.Eq("valid_description")).
					Times(1).
					Return(fakeCategory, nil)
				SUT := SaveDbCategoryService{
					category: ctgRepository,
				}
				result, err := SUT.Save("valid_name", "valid_description")
				require.NoError(t, err)
				require.Equal(t, result, fakeCategory)
			},
		},
		{
			name: "Should throw error if exists",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					Save(gomock.Any(), gomock.Any()).
					Times(1).
					Return(models.Category{}, errors.New("fake_error"))
				SUT := SaveDbCategoryService{
					category: ctgRepository,
				}
				_, err := SUT.Save("valid_name", "valid_description")
				require.NotEmpty(t, err)
				require.Equal(t, err.Error(), "fake_error")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			tc.testCase(t, ctrl)
		})
	}
}
