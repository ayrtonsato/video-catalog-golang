package routes

import (
	"database/sql"
	"net/http"

	"github.com/ayrtonsato/video-catalog-golang/internal/controllers"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	router *gin.Engine
	db     *sql.DB
}

func NewCategoryRoutes(router *gin.Engine, db *sql.DB) CategoryRoutes {
	return CategoryRoutes{
		router, db,
	}
}

func (r *CategoryRoutes) routes() {
	r.router.POST("/category", r.getCategories)
	r.router.GET("/category", r.createCategory)
}

func (r *CategoryRoutes) getCategories(ctx *gin.Context) {
	repository := repositories.NewCategoryRepository(r.db)
	service := services.NewGetCategoriesDbService(&repository)
	controller := controllers.NewGetCategoriesController(&service)
	resp := controller.Handle()
	ctx.JSON(resp.Code, gin.H{
		"body":  resp.Body,
		"error": resp.Error,
	})
}

func (r *CategoryRoutes) createCategory(ctx *gin.Context) {
	var json controllers.SaveCategoryDTO
	if err := ctx.ShouldBindJSON(&json); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
		return
	}
	validation := controllers.NewSaveCategoryValidation(&json)
	repository := repositories.NewCategoryRepository(r.db)
	service := services.NewSaveDbCategoryService(&repository)
	controller := controllers.NewSaveCategoryController(&service, json, validation)
	resp := controller.Handle()
	ctx.JSON(resp.Code, gin.H{
		"body":  resp.Body,
		"error": resp.Error,
	})
}
