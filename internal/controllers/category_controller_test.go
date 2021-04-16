package controllers

import (
	"errors"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
	"testing"
	"time"

	mock_protocols "github.com/ayrtonsato/video-catalog-golang/internal/protocols/mocks"
	mock_services "github.com/ayrtonsato/video-catalog-golang/internal/services/mocks"
	"github.com/gofrs/uuid"

	"github.com/ayrtonsato/video-catalog-golang/internal/helpers"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"

	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/stretchr/testify/require"
)

func TestGetCategoriesController_Handle(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	testCategory := []models.Category{
		{
			Id:          uid,
			Name:        "valid_category",
			Description: "valid_description",
			IsActive:    true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			DeletedAt:   nil,
		},
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should call Get repository",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				getCategories := mock_services.NewMockReaderCategory(ctrl)
				getCategories.EXPECT().GetCategories().Times(1)
				SUT := &GetCategoriesController{
					category: getCategories,
				}
				SUT.Handle()
			},
		},
		{
			name: "Should return 200 with list of categories",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				getCategories := mock_services.NewMockReaderCategory(ctrl)
				getCategories.EXPECT().GetCategories().Return(testCategory, nil)
				SUT := &GetCategoriesController{
					category: getCategories,
				}
				result := SUT.Handle()
				isEqual := cmp.Equal(testCategory, result.Body.([]models.Category))
				require.Equal(t, result.Code, 200)
				require.True(t, isEqual)
			},
		},
		{
			name: "Should return 500 when SUT throws an error",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				getCategories := mock_services.NewMockReaderCategory(ctrl)
				getCategories.EXPECT().GetCategories().Return([]models.Category{}, errors.New("new error"))
				SUT := &GetCategoriesController{
					category: getCategories,
				}
				result := SUT.Handle()
				require.Equal(t, result, helpers.HTTPInternalError())
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

func TestSaveCategoryController_Handle(t *testing.T) {
	validDTO := SaveCategoryDTO{Name: "valid_name", Description: "valid_description"}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should call validate and save correctly",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				saveCategories := mock_services.NewMockWriterCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Times(1).Return(nil)
				saveCategories.EXPECT().Save(gomock.Any(), gomock.Any()).Times(1)
				SUT := &SaveCategoryController{
					category:   saveCategories,
					validation: nameValidationMock,
					dto:        validDTO,
				}
				SUT.Handle()
			},
		},
		{
			name: "Should return 404 when validation fails",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				saveCategories := mock_services.NewMockWriterCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Times(1).Return(errors.New("invalid field"))
				SUT := &SaveCategoryController{
					category:   saveCategories,
					validation: nameValidationMock,
					dto:        validDTO,
				}
				response := SUT.Handle()
				require.Equal(t, response.Body.(error).Error(), "invalid field")
				require.Equal(t, response.Code, 400)
			},
		},
		{
			name: "Should call category service save method correctly",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				saveCategories := mock_services.NewMockWriterCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Return(nil)
				saveCategories.
					EXPECT().
					Save(gomock.Eq("valid_name"), gomock.Eq("valid_description")).
					Times(1)
				SUT := &SaveCategoryController{
					category:   saveCategories,
					validation: nameValidationMock,
					dto:        validDTO,
				}
				SUT.Handle()
			},
		},
		{
			name: "Should return 500 if service throws",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				saveCategories := mock_services.NewMockWriterCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Return(nil)
				saveCategories.
					EXPECT().
					Save(gomock.Any(), gomock.Any()).
					Times(1).
					Return(models.Category{}, errors.New("any error"))
				SUT := &SaveCategoryController{
					category:   saveCategories,
					validation: nameValidationMock,
					dto:        SaveCategoryDTO{Name: "valid_name", Description: "valid_description"},
				}
				response := SUT.Handle()
				require.Equal(t, response.Code, 500)
			},
		},
		{
			name: "Should return 200 when service category returns a category",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				uid, err := uuid.NewV4()
				if err != nil {
					t.Fatal(err)
				}
				newCategoryFake := models.Category{
					Id:          uid,
					Name:        "valid_name",
					Description: "valid_description",
					DeletedAt:   nil,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
					IsActive:    true,
				}
				saveCategories := mock_services.NewMockWriterCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Return(nil)
				saveCategories.
					EXPECT().
					Save(gomock.Eq("valid_name"), gomock.Eq("valid_description")).
					Times(1).
					Return(newCategoryFake, nil)
				SUT := &SaveCategoryController{
					category:   saveCategories,
					validation: nameValidationMock,
					dto:        validDTO,
				}
				response := SUT.Handle()
				require.Equal(t, response.Code, 201)
				require.Equal(t, response.Body, newCategoryFake)

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

func TestUpdateCategoryController_Handle(t *testing.T) {
	fakeParams := make(map[string]interface{}, 0)
	newUUID := uuid.Must(uuid.NewV4())
	fakeParams["id"] = newUUID
	validDTO := UpdateCategoryDTO{
		Name:        "valid_name",
		Description: "valid_description",
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should return 204 No Content",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				updateServicesCategory := mock_services.NewMockUpdateCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Times(1).Return(nil)
				updateServicesCategory.
					EXPECT().
					Update(
						gomock.Eq(newUUID),
						gomock.Eq(validDTO.Name),
						gomock.Eq(validDTO.Description)).
					Times(1)
				SUT := &UpdateCategoryController{
					category:   updateServicesCategory,
					validation: nameValidationMock,
					dto:        validDTO,
					params:     fakeParams,
				}
				resp := SUT.Handle()
				require.Equal(t, resp, helpers.HTTPOkNoContent())
			},
		},
		{
			name: "Should return 404 when validation fails",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				updateServicesCategory := mock_services.NewMockUpdateCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Times(1).Return(errors.New("invalid field"))
				SUT := &UpdateCategoryController{
					category:   updateServicesCategory,
					validation: nameValidationMock,
					dto:        validDTO,
					params:     fakeParams,
				}
				response := SUT.Handle()
				require.Equal(t, response.Body.(error).Error(), "invalid field")
				require.Equal(t, response.Code, 400)
			},
		},
		{
			name: "Should return 500 if service throws error",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				updateServicesCategory := mock_services.NewMockUpdateCategory(ctrl)
				nameValidationMock := mock_protocols.NewMockValidation(ctrl)
				nameValidationMock.EXPECT().Validate().Times(1).Return(nil)
				updateServicesCategory.
					EXPECT().
					Update(
						gomock.Eq(newUUID),
						gomock.Eq(validDTO.Name),
						gomock.Eq(validDTO.Description)).
					Times(1).
					Return(errors.New("any error"))
				SUT := &UpdateCategoryController{
					category:   updateServicesCategory,
					validation: nameValidationMock,
					dto:        validDTO,
					params:     fakeParams,
				}
				response := SUT.Handle()
				require.Equal(t, response.Code, 500)
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

func TestDeleteCategoryController_Handle(t *testing.T) {
	fakeParams := make(map[string]interface{}, 0)
	newUUID := uuid.Must(uuid.NewV4())
	fakeParams["id"] = newUUID
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should return 204 No Content",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				deleteServicesCategory := mock_services.NewMockDeleteCategory(ctrl)
				deleteServicesCategory.
					EXPECT().
					Delete(
						gomock.Eq(newUUID)).
					Times(1)
				SUT := NewDeleteCategoryController(deleteServicesCategory, fakeParams)
				resp := SUT.Handle()
				require.Equal(t, resp, helpers.HTTPOkNoContent())
			},
		},
		{
			name: "Should return 404 when uuid is nil",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				deleteServicesCategory := mock_services.NewMockDeleteCategory(ctrl)
				wrongFakeParam := map[string]interface{}{
					"id": uuid.Nil,
				}
				SUT := NewDeleteCategoryController(deleteServicesCategory, wrongFakeParam)
				response := SUT.Handle()
				require.Equal(t, response.Code, 404)
			},
		},
		{
			name: "Should return 404 when uuid not found",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				deleteServicesCategory := mock_services.NewMockDeleteCategory(ctrl)
				deleteServicesCategory.
					EXPECT().
					Delete(gomock.Eq(fakeParams["id"])).
					Times(1).
					Return(services.ErrCategoryNotFound)
				SUT := NewDeleteCategoryController(deleteServicesCategory, fakeParams)
				response := SUT.Handle()
				require.Equal(t, response.Code, 404)
			},
		},
		{
			name: "Should return 500 when delete services throws error",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				deleteServicesCategory := mock_services.NewMockDeleteCategory(ctrl)
				deleteServicesCategory.
					EXPECT().
					Delete(gomock.Eq(fakeParams["id"])).
					Times(1).
					Return(services.ErrCategoryUpdate)
				SUT := NewDeleteCategoryController(deleteServicesCategory, fakeParams)
				response := SUT.Handle()
				require.Equal(t, response.Code, 500)
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
