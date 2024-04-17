package repository

import (
	model_item_list "dizeto-backend/app/model/item"
	model "dizeto-backend/app/model/pricing"

	"github.com/jinzhu/gorm"
)

type PricingRepository interface {
	CreatePricing(pricing *model.Pricing, itemList []*model_item_list.ItemList) error
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

func (pr *pricingRepository) CreatePricing(pricing *model.Pricing, itemList []*model_item_list.ItemList) error {
	if err := pricing.Validate(); err != nil {
		return err
	}
	// Membuka transaksi database
	tx := pr.db.Begin()
	defer func() {
		// Jika terjadi kesalahan, rollback transaksi
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Menyimpan objek Pricing
	if err := tx.Create(pricing).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Menyimpan setiap objek ItemList secara terpisah
	for _, item := range itemList {
		if err := tx.Create(item).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	// Commit transaksi jika semua operasi berhasil
	return tx.Commit().Error
}

func (pr *pricingRepository) GetAllPricing() ([]*model.Pricing, error) {
	var pricings []*model.Pricing
	if err := pr.db.Preload("ItemList").Find(&pricings).Error; err != nil {
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
