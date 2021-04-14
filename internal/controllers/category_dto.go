package controllers

import "github.com/gofrs/uuid"

type SaveCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryDTO struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
