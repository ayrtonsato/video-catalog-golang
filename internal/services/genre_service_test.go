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

func TestGetGenresDBService_GetGenres(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	var fakeGenre = models.Genre{
		ID:        uid,
		Name:      "valid_genre",
		IsActive:  true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should return an slice of models genres without errors",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				listGenres := []models.Genre{fakeGenre}
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetGenres().
					Times(1).
					Return(listGenres, nil)
				SUT := NewGetGenresDBService(genreRepo)
				result, err := SUT.GetGenres()
				require.NoError(t, err)
				require.Equal(t, result, listGenres)
			},
		},
		{
			name: "Should throw error if exists",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetGenres().
					Times(1).
					Return([]models.Genre{}, errors.New("fake_error"))
				SUT := NewGetGenresDBService(genreRepo)
				_, err := SUT.GetGenres()
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

func TestGetGenresDBService_GetGenreByID(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	fakeGenre := models.Genre{
		ID:        uid,
		Name:      "valid_name",
		IsActive:  true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should get genre",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(uid).
					Times(1).
					Return(fakeGenre, nil)
				SUT := NewGetGenresDBService(genreRepo)
				result, err := SUT.GetGenreByID(uid)
				require.NoError(t, err)
				require.Equal(t, result, fakeGenre)
			},
		},
		{
			name: "Should return ErrNotFound",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(uid).
					Times(1).
					Return(models.Genre{}, repositories.ErrNoResult)
				SUT := NewGetGenresDBService(genreRepo)
				result, err := SUT.GetGenreByID(uid)
				require.True(t, err.Error() == ErrNotFound.Error())
				require.Equal(t, result, models.Genre{})
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

func TestGetGenresDBService_Save(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}
	listCategories := []uuid.UUID{uuid.Must(uuid.NewV4()), uuid.Must(uuid.NewV4())}
	var fakeGenre = models.Genre{
		ID:        uid,
		Name:      "valid_name",
		IsActive:  true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should return a model category without errors",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					Save(gomock.Eq(fakeGenre.Name), gomock.Eq(listCategories)).
					Times(1).
					Return(fakeGenre, nil)
				SUT := NewSaveGenreDBService(genreRepo)
				result, err := SUT.Save("valid_name", listCategories)
				require.NoError(t, err)
				require.Equal(t, result, fakeGenre)
			},
		},
		{
			name: "Should throw error if exists",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					Save(gomock.Eq(fakeGenre.Name), gomock.Eq(listCategories)).
					Times(1).
					Return(models.Genre{}, errors.New("fake_error"))
				SUT := NewSaveGenreDBService(genreRepo)
				_, err := SUT.Save("valid_name", listCategories)
				require.NotEmpty(t, err)
				require.ErrorIs(t, err, ErrSaveFailed)
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

func TestGetGenresDBService_Update(t *testing.T) {
	uid := uuid.Must(uuid.NewV4())
	fakeName := "fake_name"
	fields := []string{"name"}
	var fakeGenre = models.Genre{
		ID:        uid,
		Name:      "valid_name",
		IsActive:  true,
		DeletedAt: nil,
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Should update genre",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(uid)).
					Times(1).
					Return(fakeGenre, nil)
				genreRepo.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq(fields),
						gomock.Eq(fakeName)).
					Times(1)
				SUT := NewUpdateGenreDBService(genreRepo)
				err := SUT.Update(uid, "fake_name")
				require.NoError(t, err)
			},
		},
		{
			name: "Should throw error when update genre fails",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(uid)).
					Times(1).
					Return(fakeGenre, nil)
				genreRepo.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq(fields),
						gomock.Eq(fakeName)).
					Times(1).Return(ErrUpdateFailed)
				SUT := NewUpdateGenreDBService(genreRepo)
				err := SUT.Update(uid, "fake_name")
				require.Error(t, err)
				require.True(t, err.Error() == ErrUpdateFailed.Error())
			},
		},
		{
			name: "Should throw error genre not found",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(uid)).
					Times(1).
					Return(models.Genre{}, ErrNotFound)
				genreRepo.
					EXPECT().
					Update(
						gomock.Eq(uid),
						gomock.Eq(fields),
						gomock.Eq(fakeName)).
					Times(0)
				SUT := NewUpdateGenreDBService(genreRepo)
				err := SUT.Update(uid, "fake_name")
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

func TestDeleteGenreDBService_Delete(t *testing.T) {
	fakeGenre := models.Genre{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      "fake_name",
		IsActive:  true,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, ctrl *gomock.Controller)
	}{
		{
			name: "Delete genre successfully",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(fakeGenre.ID)).
					Times(1).Return(fakeGenre, nil)
				genreRepo.
					EXPECT().
					Delete(gomock.Eq(fakeGenre)).
					Times(1).Return(nil)
				SUT := NewDeleteGenreDBService(genreRepo)
				err := SUT.Delete(fakeGenre.ID)
				require.NoError(t, err)
			},
		},
		{
			name: "Throw ErrNoResult when no genre found",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(fakeGenre.ID)).
					Times(1).Return(models.Genre{}, repositories.ErrNoResult)
				genreRepo.
					EXPECT().
					Delete(gomock.Eq(fakeGenre.ID)).
					Times(0)
				SUT := NewDeleteGenreDBService(genreRepo)
				err := SUT.Delete(fakeGenre.ID)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNotFound)
			},
		},
		{
			name: "Throw ErrUpdateFailed when update genre failed",
			testCase: func(t *testing.T, ctrl *gomock.Controller) {
				genreRepo := mock_repositories.NewMockGenreDB(ctrl)
				genreRepo.
					EXPECT().
					GetByID(gomock.Eq(fakeGenre.ID)).
					Times(1).Return(fakeGenre, nil)
				genreRepo.
					EXPECT().
					Delete(gomock.Eq(fakeGenre)).
					Times(1).Return(repositories.ErrOnUpdate)
				SUT := NewDeleteGenreDBService(genreRepo)
				err := SUT.Delete(fakeGenre.ID)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrUpdateFailed)
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
