package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Video struct {
	Id           uuid.UUID  `json:"id"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	YearLaunched int        `json:"year_launched"`
	Opened       bool       `json:"opened"`
	Rating       string     `json:"rating"`
	Duration     int        `json:"duration"`
	IsActive     bool       `json:"isActive"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}

func NewVideo() Video {
	return Video{}
}
