package repository

import (
	model "dizeto-backend/app/model/page"

	"github.com/jinzhu/gorm"
)

type LandingRepository interface {
	GetLandingPage() ([]*model.Page, error)
}

type landingRepository struct {
	db *gorm.DB
}

func NewLandingRepository(db *gorm.DB) LandingRepository {
	return &landingRepository{db: db}
}

func (lr *landingRepository) GetLandingPage() ([]*model.Page, error) {
	var pages []*model.Page
	if err := lr.db.
		Preload("Abouts").
		Preload("Clients").
		Preload("Countings").
		Preload("Highlights").
		Preload("Pricings", func(db *gorm.DB) *gorm.DB {
			return db.Preload("ItemList")
		}).
		Preload("Testimonis").
		Find(&pages).Error; err != nil {
		return nil, err
	}
	return pages, nil
}
