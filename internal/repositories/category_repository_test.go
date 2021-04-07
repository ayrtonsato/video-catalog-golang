package repositories

import (
	"database/sql"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/gofrs/uuid"
	"github.com/google/go-cmp/cmp"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestCategoryRepository_GetCategories(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	var fakeCategory = models.Category{
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock)
	}{
		{
			name: "Should return an array of categories",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) {
				fakeListCategory := []models.Category{
					fakeCategory,
				}
				SUT := CategoryRepository{
					db: db,
				}
				re := regexp.
					QuoteMeta("SELECT id, name, description, is_active, created_at, updated_at, deleted_at FROM categories")

				mock.ExpectQuery(re).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "name", "description", "is_active", "created_at", "updated_at", "deleted_at",
						}).AddRow(
							fakeCategory.Id,
							fakeCategory.Name,
							fakeCategory.Description,
							true,
							fakeCategory.CreatedAt,
							fakeCategory.UpdatedAt, nil))
				list, err := SUT.GetCategories()
				require.NoError(t, err)
				require.True(t, cmp.Equal(list, fakeListCategory))

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Should return an empty array of categories with ErrNoRows",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) {
				SUT := CategoryRepository{
					db: db,
				}
				re := regexp.
					QuoteMeta("SELECT id, name, description, is_active, created_at, updated_at, deleted_at FROM categories")

				mock.ExpectQuery(re).
					WillReturnError(sql.ErrNoRows)
				list, err := SUT.GetCategories()
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNoRows)

				require.True(t, len(list) == 0)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tc.testCase(t, db, mock)
		})
	}
}

func TestCategoryRepository_Save(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	var fakeCategory = models.Category{
		Id:          uid,
		Name:        "valid_name",
		Description: "valid_description",
		IsActive:    true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   nil,
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock)
	}{
		{
			name: "Should save new Category and return it",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) {
				SUT := CategoryRepository{
					db: db,
				}
				expectStmt := mock.ExpectPrepare("^INSERT INTO categories.*")
				rows := sqlmock.
					NewRows([]string{"id", "name", "description", "is_active", "created_at", "updated_at", "deleted_at"}).
					AddRow(fakeCategory.Id,
						fakeCategory.Name,
						fakeCategory.Description,
						fakeCategory.IsActive,
						fakeCategory.CreatedAt,
						fakeCategory.UpdatedAt,
						fakeCategory.DeletedAt)
				expectStmt.
					ExpectQuery().
					WithArgs("valid_name", "valid_description").
					WillReturnRows(rows)
				category, err := SUT.Save("valid_name", "valid_description")

				require.NoError(t, err)
				require.Equal(t, category, fakeCategory)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Should return an empty object with ErrNoRows",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) {
				SUT := CategoryRepository{
					db: db,
				}
				expectStmt := mock.ExpectPrepare("^INSERT INTO categories.*")
				expectStmt.
					ExpectQuery().
					WithArgs("invalid_name", "invalid_description").
					WillReturnError(sql.ErrNoRows)
				category, err := SUT.Save("invalid_name", "invalid_description")
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNoRows)
				require.Equal(t, category, models.Category{})
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tc.testCase(t, db, mock)
		})
	}
}

/*func TestCategoryRepository_GetCategoryByUUID(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock)
	}{
		{
			name: "Should return an array of categories",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock) {
				newUUID := uuid.New()
				createdAt := time.Now()
				updatedAt := time.Now()

				fakeListCategory := []models.Category{{
					Id:          newUUID,
					Name:        "test_category",
					Description: "description_category",
					IsActive:    true,
					CreatedAt:   createdAt,
					UpdatedAt:   updatedAt,
					DeletedAt:   nil,
				}}

				SUT := CategoryRepository{
					db: db,
				}

				mock.ExpectQuery(regexp.
					QuoteMeta("SELECT id, name, description, is_active, created_at, updated_at, deleted_at FROM categories")).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "name", "description", "is_active", "created_at", "updated_at", "deleted_at",
						}).AddRow(
							newUUID, "test_category", "description_category", true, createdAt, updatedAt, nil))
				list, err := SUT.GetCategories()
				require.NoError(t, err)
				require.True(t, cmp.Equal(list, fakeListCategory))

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tc.testCase(t, db, mock)
		})
	}
}*/
