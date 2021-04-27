package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	mock_logger "github.com/ayrtonsato/video-catalog-golang/pkg/logger/mocks"
	"github.com/gofrs/uuid"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
	"time"
)

func TestGenreRepository_GetGenres(t *testing.T) {
	uid, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	var fakeGenre = models.Genre{
		ID:        uid,
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
			name: "Return an array of genres",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				fakeListGenre := []models.Genre{
					fakeGenre,
				}
				SUT := NewGenreRepository(db, log)
				re := regexp.
					QuoteMeta("SELECT id, name, is_active, created_at, updated_at, deleted_at FROM genres")

				mock.ExpectQuery(re).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "name", "is_active", "created_at", "updated_at", "deleted_at",
						}).AddRow(
							fakeGenre.ID,
							fakeGenre.Name,
							true,
							fakeGenre.CreatedAt,
							fakeGenre.UpdatedAt, nil))
				list, err := SUT.GetGenres()
				require.NoError(t, err)
				require.True(t, cmp.Equal(list, fakeListGenre))

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Return ErrNoRows error",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				log.EXPECT().Error(sql.ErrNoRows.Error()).Times(1)
				SUT := NewGenreRepository(db, log)
				re := regexp.
					QuoteMeta("SELECT id, name, is_active, created_at, updated_at, deleted_at FROM genres")

				mock.ExpectQuery(re).
					WillReturnError(sql.ErrNoRows)
				_, err := SUT.GetGenres()
				require.Error(t, err)
				require.ErrorIs(t, err, ErrNoResult)

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
	var fakeGenre = models.Genre{
		ID:        uuid.Must(uuid.NewV4()),
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
			name: "Return genre successfully",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				fields := sqlmock.NewRows([]string{
					"id", "name", "is_active",
					"created_at", "updated_at", "deleted_at",
				}).AddRow(fakeGenre.ID, fakeGenre.Name,
					fakeGenre.IsActive, fakeGenre.CreatedAt,
					fakeGenre.UpdatedAt, fakeGenre.DeletedAt)
				log := mock_logger.NewMockLogger(ctrl)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				re := regexp.QuoteMeta("SELECT * FROM genres WHERE id=$1")
				mock.ExpectQuery(re).
					WithArgs(fakeGenre.ID).WillReturnRows(fields)
				genre, err := SUT.GetGenreByID(fakeGenre.ID)
				require.NoError(t, err)
				require.Equal(t, genre, fakeGenre)

				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Throw ErrNoResult error",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				log.EXPECT().Error(sql.ErrNoRows.Error()).Times(1)
				SUT := NewGenreRepository(db, log)
				re := regexp.QuoteMeta("SELECT * FROM genres WHERE id=$1")
				mock.ExpectQuery(re).
					WillReturnError(sql.ErrNoRows)
				_, err := SUT.GetGenreByID(fakeGenre.ID)

				require.Error(t, err)
				require.ErrorIs(t, err, ErrNoResult)

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

func TestGenreRepository_Save(t *testing.T) {
	catID := []uuid.UUID{uuid.Must(uuid.NewV4())}
	uid := uuid.Must(uuid.NewV4())

	var fakeGenre = models.Genre{
		ID:        uid,
		Name:      "valid_name",
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
			name: "Save new genre and return it",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := NewGenreRepository(db, log)
				mock.ExpectBegin()
				expectStmt := mock.ExpectPrepare("^INSERT INTO genres.*")
				rows := sqlmock.
					NewRows([]string{"id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
					AddRow(fakeGenre.ID,
						fakeGenre.Name,
						fakeGenre.IsActive,
						fakeGenre.CreatedAt,
						fakeGenre.UpdatedAt,
						fakeGenre.DeletedAt)
				expectStmt.
					ExpectQuery().
					WithArgs("valid_name").
					WillReturnRows(rows)
				expectRelationStmt := mock.ExpectPrepare("^INSERT INTO categories_genres.*")
				expectRelationStmt.
					ExpectExec().
					WithArgs(catID[0], fakeGenre.ID).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
				genre, err := SUT.Save("valid_name", catID)
				require.NoError(t, err)
				require.Equal(t, genre, fakeGenre)
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Return ErrOnSave when insert into genres failed",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := NewGenreRepository(db, log)
				log.EXPECT().Error(sql.ErrNoRows.Error()).Times(1)
				mock.ExpectBegin()
				expectStmt := mock.ExpectPrepare("^INSERT INTO genres.*")
				expectStmt.
					ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
					WillReturnError(sql.ErrNoRows)
				mock.ExpectRollback()
				_, err := SUT.Save("invalid_name", catID)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrOnSave)
				if err := mock.ExpectationsWereMet(); err != nil {
					t.Errorf("%s", err)
				}
			},
		},
		{
			name: "Return ErrOnSave when insert into categories_genres failed",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := NewGenreRepository(db, log)
				mock.ExpectBegin()
				expectStmt := mock.ExpectPrepare("^INSERT INTO genres.*")
				log.EXPECT().Error(sql.ErrNoRows.Error()).Times(1)
				rows := sqlmock.
					NewRows([]string{"id", "name", "is_active", "created_at", "updated_at", "deleted_at"}).
					AddRow(fakeGenre.ID,
						fakeGenre.Name,
						fakeGenre.IsActive,
						fakeGenre.CreatedAt,
						fakeGenre.UpdatedAt,
						fakeGenre.DeletedAt)
				expectStmt.
					ExpectQuery().
					WithArgs(sqlmock.AnyArg()).
					WillReturnRows(rows)
				expectRelationStmt := mock.ExpectPrepare("^INSERT INTO categories_genres.*")
				expectRelationStmt.
					ExpectExec().
					WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
					WillReturnError(sql.ErrNoRows)
				mock.ExpectRollback()

				_, err := SUT.Save("invalid_name", catID)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrOnSave)
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

func TestGenreRepository_Update(t *testing.T) {
	newUUID := uuid.Must(uuid.NewV4())
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller)
	}{
		{
			name: "Update genre successfully",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				expectStmt := mock.ExpectPrepare("^UPDATE genres SET.*")
				expectStmt.
					ExpectExec().
					WithArgs("other_name", newUUID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				log.EXPECT().Error(gomock.Any()).Times(0)
				fields := []string{"name"}
				values := []interface{}{"other_name"}
				err := SUT.Update(newUUID, fields, values...)
				require.NoError(t, err)
			},
		},
		{
			name: "Throw ErrOnUpdate error",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				log.EXPECT().Error(gomock.Eq(sql.ErrConnDone.Error())).Times(1)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				expectStmt := mock.ExpectPrepare("^UPDATE genres SET.*")
				expectStmt.
					ExpectExec().
					WithArgs("other_name", newUUID).
					WillReturnError(sql.ErrConnDone)

				fields := []string{"name"}
				values := []interface{}{"other_name"}
				err := SUT.Update(newUUID, fields, values...)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrOnUpdate)
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

func TestGenreRepository_Delete(t *testing.T) {
	// deletedTime := time.Now().UTC()
	fakeGenre := models.Genre{
		ID:        uuid.Must(uuid.NewV4()),
		Name:      "fake_genre",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		DeletedAt: nil,
		IsActive:  true,
	}
	testCases := []struct {
		name     string
		testCase func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller)
	}{
		{
			name: "Delete genre successfully",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				mock.ExpectBegin()
				expectUpdateStmt := mock.ExpectPrepare("^UPDATE genres SET.*")
				expectUpdateStmt.
					ExpectExec().
					WithArgs(false, sqlmock.AnyArg(), fakeGenre.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				re := regexp.QuoteMeta("DELETE FROM categories_genres WHERE genre_id=$1")
				expectDeleteStmt := mock.ExpectPrepare(re)
				expectDeleteStmt.
					ExpectExec().
					WithArgs(fakeGenre.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
				err := SUT.Delete(fakeGenre)
				require.NoError(t, err)
			},
		},
		{
			name: "Throw ErrOnDelete when update return an error",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				log.EXPECT().Error(sql.ErrConnDone.Error()).Times(1)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				mock.ExpectBegin()
				expectUpdateStmt := mock.ExpectPrepare("^UPDATE genres SET.*")
				expectUpdateStmt.
					ExpectExec().
					WithArgs(false, sqlmock.AnyArg(), fakeGenre.ID).
					WillReturnResult(sqlmock.NewResult(1, 1))
				re := regexp.QuoteMeta("DELETE FROM categories_genres WHERE genre_id=$1")
				expectDeleteStmt := mock.ExpectPrepare(re)
				expectDeleteStmt.
					ExpectExec().
					WithArgs(fakeGenre.ID).
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
				err := SUT.Delete(fakeGenre)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrOnDelete)
			},
		},
		{
			name: "Throw ErrOnDelete when delete categories_genres fail",
			testCase: func(t *testing.T, db *sql.DB, mock sqlmock.Sqlmock, ctrl *gomock.Controller) {
				log := mock_logger.NewMockLogger(ctrl)
				log.EXPECT().Error(sql.ErrConnDone.Error()).Times(1)
				SUT := GenreRepository{
					db:  db,
					log: log,
				}
				mock.ExpectBegin()
				expectUpdateStmt := mock.ExpectPrepare("^UPDATE genres SET.*")
				expectUpdateStmt.
					ExpectExec().
					WithArgs(false, sqlmock.AnyArg(), fakeGenre.ID).
					WillReturnError(sql.ErrConnDone)
				mock.ExpectRollback()
				err := SUT.Delete(fakeGenre)
				require.Error(t, err)
				require.ErrorIs(t, err, ErrOnDelete)
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
