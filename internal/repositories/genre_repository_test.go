package repositories

import (
	_ "github.com/jackc/pgx/v4/stdlib"
)

/*func TestGenreRepository_GetGenres(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	var fakeCategory = models.Genre{
		Id:        uid,
		Name:      "valid_genre",
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller)
	}{
		{
			name: "Should return an array of categories",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				fakeListCategory := []models.Category{
					fakeCategory,
				}
				SUT := NewCategoryRepository(db, log)
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
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := CategoryRepository{
					db:  db,
					log: log,
				}
				re := regexp.
					QuoteMeta("SELECT id, name, description, is_active, created_at, updated_at, deleted_at FROM categories")

				mock.ExpectQuery(re).
					WillReturnError(sql.ErrNoRows)
				log.EXPECT().Error(gomock.Any()).Times(1)
				list, err := SUT.GetCategories()
				require.Error(t, err)
				require.ErrorIs(t, err, sql.ErrNoRows)

				require.True(t, len(list) == 0)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tc.testCase(t, db, mock, ctrl)
		})
	}
}

func TestGenreRepository_GetGenreByID(t *testing.T) {
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller)
	}{
		{
			name: "Should return a category successfully",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				newUUID, err := uuid.NewV4()
				require.NoError(t, err)
				fakeCategory := models.Category{
					Id:          newUUID,
					Name:        "test_name",
					Description: "test_desc",
					CreatedAt:   time.Now().UTC(),
					UpdatedAt:   time.Now().UTC(),
					IsActive:    true,
					DeletedAt:   nil,
				}
				fields := sqlmock.NewRows([]string{
					"id", "name",
					"description", "is_active",
					"created_at", "updated_at", "deleted_at",
				}).AddRow(fakeCategory.Id, fakeCategory.Name, fakeCategory.Description,
					fakeCategory.IsActive, fakeCategory.CreatedAt,
					fakeCategory.UpdatedAt, fakeCategory.DeletedAt)
				log := mock_logger.NewMockLogger(ctrl)
				SUT := CategoryRepository{
					db:  db,
					log: log,
				}
				re := regexp.QuoteMeta("SELECT * FROM categories WHERE id=$1")
				mock.ExpectQuery(re).
					WithArgs(newUUID).WillReturnRows(fields)
				category, err := SUT.GetByID(newUUID)
				require.NoError(t, err)
				require.Equal(t, category, fakeCategory)
			},
		},
		{
			name: "Should throw error",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				newUUID := uuid.Must(uuid.NewV4())
				log := mock_logger.NewMockLogger(ctrl)
				SUT := CategoryRepository{
					db:  db,
					log: log,
				}
				re := regexp.QuoteMeta("SELECT * FROM categories WHERE id=$1")
				mock.ExpectQuery(re).
					WithArgs(newUUID).
					WillReturnError(sql.ErrNoRows)
				log.EXPECT().Error(gomock.Eq(sql.ErrNoRows.Error())).Times(1)
				_, err := SUT.GetByID(newUUID)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNoResult)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tc.testCase(t, db, mock, ctrl)
		})
	}
}
*/
