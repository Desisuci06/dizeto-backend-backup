package model

import (
	"time"

	model_item_list "dizeto-backend/app/model/item"

	"github.com/go-playground/validator/v10"

	"github.com/google/uuid"
)

type Pricing struct {
	ID        uuid.UUID                   `gorm:"type:uuid;primary_key;" json:"id"`
	Title     string                      `gorm:"not null" json:"title" validate:"required,min=3,max=255"`
	Price     uint                        `gorm:"not null" json:"price" validate:"required"`
	Paket     string                      `gorm:"not null" json:"paket" validate:"required"`
	Category  string                      `gorm:"not null" json:"category"`
	ItemList  []*model_item_list.ItemList `gorm:"foreignkey:PricingID" json:"item_list"`
	PageID    uint                        `json:"page_id"`
	CreatedAt time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time                   `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Pricing) Validate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return err
	}

	return nil
}
