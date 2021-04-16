package routes

import (
	"database/sql"
	"github.com/ayrtonsato/video-catalog-golang/internal/controllers"
	"github.com/ayrtonsato/video-catalog-golang/internal/repositories"
	"github.com/ayrtonsato/video-catalog-golang/internal/services"
	"github.com/ayrtonsato/video-catalog-golang/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
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
	r.router.PUT("/category/:id", r.updateCategory)
	r.router.DELETE("/category/:id", r.deleteCategory)
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

func (r *CategoryRoutes) updateCategory(ctx *gin.Context) {
	params := make(map[string]interface{})

	id := ctx.Param("id")
	newUUID, err := uuid.FromString(id)
	if err != nil {
		r.log.Error(err)
	}

	params["id"] = newUUID

	var dto controllers.UpdateCategoryDTO
	if err = ctx.ShouldBind(&dto); err != nil {
		r.log.Error(err)
	}
	val := controllers.NewUpdateCategoryValidation(&dto)
	repo := repositories.NewCategoryRepository(r.db, r.log)
	serv := services.NewUpdateDbCategoryService(&repo)
	ctrl := controllers.NewUpdateCategoryController(&serv, dto, val, params)
	resp := ctrl.Handle()

	ctx.JSON(resp.Code, resp.Body)
}

func (r *CategoryRoutes) deleteCategory(ctx *gin.Context) {
	params := make(map[string]interface{})

	id := ctx.Param("id")
	newUUID, err := uuid.FromString(id)
	if err != nil {
		r.log.Error(err)
	}
	params["id"] = newUUID

	repo := repositories.NewCategoryRepository(r.db, r.log)
	serv := services.NewDeleteDBCategoryService(&repo)
	ctrl := controllers.NewDeleteCategoryController(&serv, params)
	resp := ctrl.Handle()

	ctx.JSON(resp.Code, resp.Body)
}
