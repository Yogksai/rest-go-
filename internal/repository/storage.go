package repository

import (
	"rest-api/internal/models"

	"gorm.io/gorm"
)

// Команды ормки будут тута типа CRUD
type UrlRepo struct {
	DB *gorm.DB
}

func NewUrlRepo(db *gorm.DB) *UrlRepo {
	return &UrlRepo{DB: db}
}

// CRUD - Create, Read, Update, Delete

func (r *UrlRepo) Create(link models.Link) error {
	err := r.DB.Create(&link).Error
	if err != nil {
		return err
	}
	return nil
}
