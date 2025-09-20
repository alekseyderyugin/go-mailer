package domain

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID        ClientId `gorm:"type:char(36);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Letters   []Letter
}

func NewClient(id ClientId) *Client {
	entity := Client{
		ID: id,
	}

	now := time.Now()
	entity.CreatedAt = now
	entity.CreatedAt = now
	return &entity
}
