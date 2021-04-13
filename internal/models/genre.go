package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type Genre struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewGenre() Genre {
	return Genre{}
}
