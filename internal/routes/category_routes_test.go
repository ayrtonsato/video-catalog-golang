package routes_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ayrtonsato/video-catalog-golang/internal/models"
	"github.com/ayrtonsato/video-catalog-golang/internal/setup"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func saveCategory(t *testing.T, db *sql.DB, name string, description string) (models.Category, error) {
	insertStatement := `INSERT INTO categories(name, description)
		VALUES($1, $2)
		RETURNING id, name, description, is_active, created_at, updated_at, deleted_at
	`
	stmt, err := db.Prepare(insertStatement)
	if err != nil {
		require.NoError(t, err)
	}
	defer stmt.Close()
	row := stmt.QueryRow(name, description)
	var category models.Category
	err = row.Scan(
		&category.Id,
		&category.Name,
		&category.Description,
		&category.IsActive,
		&category.CreatedAt,
		&category.UpdatedAt,
		&category.DeletedAt)
	if err != nil {
		require.NoError(t, err)
	}
	return category, nil
}

func deleteAllCategories(t *testing.T, db *sql.DB) error {
	deleteStatement := `DELETE FROM categories`
	_, err := db.Exec(deleteStatement)
	return err
}

func TestCategoryRoutes_GetCategories(t *testing.T) {
	testCases := []struct {
		name     string
		setup    func(t *testing.T, db *sql.DB)
		response func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "200 OK",
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			url := "/category"
			request := httptest.NewRequest(http.MethodGet, url, nil)
			tSetup := setup.TestSetup{}
			tSetup.
				BuildConfig(t, "../../").
				BuildLogger(t).
				BuildDB(t, nil).
				BuildServer(t).
				Serve(recorder, request)
			tc.response(t, recorder)
		})
	}
}

func TestCategoryRoutes_CreateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		body     gin.H
		response func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "201 Created",
			body: gin.H{
				"name":        "valid_name",
				"description": "valid_description",
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusCreated, r.Code)
			},
		},
		{
			name: "400 BadRequest without name",
			body: gin.H{
				"description": "valid_description",
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
		{
			name: "400 BadRequest without description",
			body: gin.H{
				"name": "valid_name",
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
		{
			name: "400 BadRequest invalid name",
			body: gin.H{
				"name":        "va",
				"description": "valid_desc",
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				fmt.Println(r.Body)
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			url := "/category"

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			tSetup := setup.TestSetup{}
			tSetup.
				BuildConfig(t, "../../").
				BuildLogger(t).
				BuildDB(t, nil).
				BuildServer(t).
				Serve(recorder, request)
			tc.response(t, recorder)
		})
	}
}

func TestCategoryRoutes_GetSingleCategory(t *testing.T) {
	testCases := []struct {
		name     string
		urlFn    func(category *models.Category) string
		response func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "200 OK",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, r.Code)
			},
		},
		{
			name: "404 NotFound when id not found",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", uuid.Must(uuid.NewV4()).String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
		{
			name: "404 NotFound when id is not an uuid",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/teste")
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			tSetup := setup.TestSetup{}
			tSetup.
				BuildConfig(t, "../../").
				BuildLogger(t).
				BuildDB(t, nil).
				BuildServer(t)

			category, err := saveCategory(t, tSetup.DB, "teste", "teste")
			url := tc.urlFn(&category)
			request := httptest.NewRequest(http.MethodGet, url, nil)

			tSetup.Serve(recorder, request)
			require.NoError(t, err)
			tc.response(t, recorder)
		})
	}
}

func TestCategoryRoutes_UpdateCategory(t *testing.T) {
	testCases := []struct {
		name     string
		body     gin.H
		urlFn    func(category *models.Category) string
		response func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "204 NoContent",
			body: gin.H{
				"name":        "diff_name",
				"description": "diff_desc",
			},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, r.Code)
			},
		},
		{
			name: "404 NotFound when id not found",
			body: gin.H{
				"name":        "diff_name",
				"description": "diff_desc",
			},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", uuid.Must(uuid.NewV4()).String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
		{
			name: "404 NotFound when id is not an uuid",
			body: gin.H{
				"name":        "diff_name",
				"description": "diff_desc",
			},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/teste")
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
		{
			name: "400 BadRequest when no name",
			body: gin.H{
				"description": "diff_desc",
			},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
		{
			name: "400 BadRequest when no description",
			body: gin.H{
				"name": "diff_name",
			},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
		{
			name: "400 BadRequest when no body",
			body: gin.H{},
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			tSetup := setup.TestSetup{}
			tSetup.
				BuildConfig(t, "../../").
				BuildLogger(t).
				BuildDB(t, nil).
				BuildServer(t)

			category, err := saveCategory(t, tSetup.DB, "teste", "teste")
			url := tc.urlFn(&category)

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			request := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(data))

			tSetup.Serve(recorder, request)
			require.NoError(t, err)
			tc.response(t, recorder)
		})
	}
}

func TestCategoryRoutes_DeleteCategory(t *testing.T) {
	testCases := []struct {
		name     string
		urlFn    func(category *models.Category) string
		response func(t *testing.T, recoder *httptest.ResponseRecorder)
	}{
		{
			name: "204 NoContent",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", category.Id.String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNoContent, r.Code)
			},
		},
		{
			name: "404 NotFound when id not found",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/%v", uuid.Must(uuid.NewV4()).String())
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
		{
			name: "404 NotFound when id is not an uuid",
			urlFn: func(category *models.Category) string {
				return fmt.Sprintf("/category/teste")
			},
			response: func(t *testing.T, r *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, r.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			tSetup := setup.TestSetup{}
			tSetup.
				BuildConfig(t, "../../").
				BuildLogger(t).
				BuildDB(t, nil).
				BuildServer(t)

			category, err := saveCategory(t, tSetup.DB, "teste", "teste")
			url := tc.urlFn(&category)

			request := httptest.NewRequest(http.MethodDelete, url, nil)

			tSetup.Serve(recorder, request)
			require.NoError(t, err)
			tc.response(t, recorder)
		})
	}
}
