package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type CastMember struct {
	Id        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Type      string     `json:"description"`
	IsActive  bool       `json:"isActive"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

func NewCastMember() CastMember {
	return CastMember{}
}
