package routes

import (
	"database/sql"
	"github.com/ayrtonsato/video-catalog-golang/internal/controllers"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
	"github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gin-gonic/gin"
)

type CategoryRoutes struct {
	router *gin.Engine
	db     *sql.DB
	log    logger.Logger
}

func NewCategoryRoutes(router *gin.Engine, db *sql.DB, log logger.Logger) CategoryRoutes {
	return CategoryRoutes{
		router, db, log,
	}
}

func (r CategoryRoutes) Routes() {
	r.router.POST("/category", r.createCategory)
	r.router.GET("/category", r.getCategories)
}

func (r *CategoryRoutes) getCategories(ctx *gin.Context) {
	repository := repositories.NewCategoryRepository(r.db, r.log)
	service := services.NewGetCategoriesDbService(&repository)
	controller := controllers.NewGetCategoriesController(&service)
	resp := controller.Handle()
	ctx.JSON(resp.Code, gin.H{
		"body": resp.Body,
	})
}

func (r *CategoryRoutes) createCategory(ctx *gin.Context) {
	var json controllers.SaveCategoryDTO
	if err := ctx.ShouldBindJSON(&json); err != nil {
		json = controllers.SaveCategoryDTO{}
	}
	validation := controllers.NewSaveCategoryValidation(&json)
	repository := repositories.NewCategoryRepository(r.db, r.log)
	service := services.NewSaveDbCategoryService(&repository)
	controller := controllers.NewSaveCategoryController(&service, json, validation)
	resp := controller.Handle()
	ctx.JSON(resp.Code, gin.H{
		"body": resp.Body,
	})
}
