package repository

import (
	model "dizeto-backend/app/model/pricing"

	"github.com/jinzhu/gorm"
)

type PricingRepository interface {
	CreatePricing(pricing *model.Pricing) error
	GetAllPricing() ([]*model.Pricing, error)
	GetPricingByID(id string) (*model.Pricing, error)
	UpdatePricing(pricing *model.Pricing) error
}

type pricingRepository struct {
	db *gorm.DB
}

func NewPricingRepository(db *gorm.DB) PricingRepository {
	return &pricingRepository{db: db}
}

func (pr *pricingRepository) CreatePricing(pricing *model.Pricing) error {
	if err := pricing.Validate(); err != nil {
		return err
	}
	err := pr.db.Create(pricing).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *pricingRepository) GetAllPricing() ([]*model.Pricing, error) {
	var pricings []*model.Pricing
	if err := pr.db.Find(&pricings).Error; err != nil {
		return nil, err
	}
	return pricings, nil
}

func (pr *pricingRepository) GetPricingByID(id string) (*model.Pricing, error) {
	var pricing model.Pricing
	err := pr.db.Where("id = ?", id).First(&pricing).Error
	return &pricing, err
}

func (pr *pricingRepository) UpdatePricing(pricing *model.Pricing) error {

	if err := pricing.Validate(); err != nil {
		return err
	}

	err := pr.db.Save(pricing).Error
	return err
}
