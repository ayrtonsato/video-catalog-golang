package models

import (
	"github.com/gofrs/uuid"
	"time"
)

type CastMemberType int

const (
	ACTOR CastMemberType = iota + 1
	DIRECTOR
)

type CastMember struct {
	Id        uuid.UUID      `json:"id"`
	Name      string         `json:"name"`
	Type      CastMemberType `json:"description"`
	IsActive  bool           `json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt *time.Time     `json:"deletedAt"`
}

func NewCastMember() CastMember {
	return CastMember{}
}
