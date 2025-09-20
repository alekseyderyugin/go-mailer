package infrastructure

import (
	"go-mailer/letters/domain"
	"gorm.io/gorm"
)

type ClientRepository struct {
	db      *gorm.DB
	context *Context
}

func NewClientRepository(db *gorm.DB, ctx *Context) *ClientRepository {
	return &ClientRepository{
		db:      db,
		context: ctx,
	}
}

func (rep *ClientRepository) AutoMigrate() error {
	return (*rep).db.AutoMigrate(&domain.Client{})
}

func (rep *ClientRepository) Save(client *domain.Client) error {
	err := rep.db.Save(client).Error
	return err
}
