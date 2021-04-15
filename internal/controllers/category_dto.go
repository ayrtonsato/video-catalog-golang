package controllers

type SaveCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
