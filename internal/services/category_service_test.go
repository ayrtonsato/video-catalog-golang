package services

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	mock_repositories "github.com/ayrtonsato/video-catalog-golang/internal/repositories/mocks"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
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
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		DeletedAt:   nil,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
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
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		DeletedAt:   nil,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
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

func TestUpdateDbCategoryService_Update(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	fakeName := "fake_name"
	fakeDesc := "fake_desc"
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should update category",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, nil)
				ctgRepository.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq([]string{"name", "description"}),
						gomock.Eq(fakeName),
						gomock.Eq(fakeDesc)).
					Times(1)
				SUT := UpdateDbCategoryService{
					category: ctgRepository,
				}
				err := SUT.Update(uid, fakeName, fakeDesc)
				require.NoError(t, err)
			},
		},
		{
			name: "Should throw error when update category fails",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, nil)
				ctgRepository.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq([]string{"name", "description"}),
						gomock.Eq(fakeName),
						gomock.Eq(fakeDesc)).
					Times(1).Return(ErrUpdateFailed)
				SUT := UpdateDbCategoryService{
					category: ctgRepository,
				}
				err := SUT.Update(uid, fakeName, fakeDesc)
				require.Error(t, err)
				require.True(t, err.Error() == ErrUpdateFailed.Error())
			},
		},
		{
			name: "Should throw error category not found",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, ErrNotFound).Times(1)
				ctgRepository.
					EXPECT().
					Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				SUT := UpdateDbCategoryService{
					category: ctgRepository,
				}
				err := SUT.Update(uid, fakeName, fakeDesc)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNotFound)
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

func TestDeleteDBCategoryService_Delete(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should delete category",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, nil)
				ctgRepository.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq([]string{"is_active", "deleted_at"}),
						gomock.Eq(false),
						gomock.Any()).
					Times(1)
				SUT := DeleteDBCategoryService{
					category: ctgRepository,
				}
				err := SUT.Delete(uid)
				require.NoError(t, err)
			},
		},
		{
			name: "Should throw error when delete category fails",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, nil)
				ctgRepository.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq([]string{"is_active", "deleted_at"}),
						gomock.Eq(false),
						gomock.Any()).
					Times(1).Return(ErrUpdateFailed)
				SUT := DeleteDBCategoryService{
					category: ctgRepository,
				}
				err := SUT.Delete(uid)
				require.Error(t, err)
				require.True(t, err.Error() == ErrUpdateFailed.Error())
			},
		},
		{
			name: "Should throw error when category not found",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, ErrNotFound).Times(1)
				ctgRepository.
					EXPECT().
					Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				SUT := DeleteDBCategoryService{
					category: ctgRepository,
				}
				err := SUT.Delete(uid)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNotFound)
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

func TestGetCategoriesDbService_GetCategory(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	fakeCategory := models.Category{
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		DeletedAt:   nil,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should get category",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(fakeCategory, nil)
				SUT := NewGetCategoriesDbService(ctgRepository)
				category, err := SUT.GetCategory(uid)
				require.NoError(t, err)
				require.Equal(t, category, fakeCategory)
			},
		},
		{
			name: "Should return ErrNotFound",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				ctgRepository := mock_repositories.NewMockCategory(ctrl)
				ctgRepository.
					EXPECT().
					GetByID(uid).
					Return(models.Category{}, repositories.ErrNoResult)
				SUT := NewGetCategoriesDbService(ctgRepository)
				category, err := SUT.GetCategory(uid)
				require.Error(t, err)
				require.True(t, err.Error() == ErrNotFound.Error())
				require.Equal(t, category, models.Category{})
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
