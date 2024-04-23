package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ItemList struct {
	ID        uuid.UUID `gorm:"primary_key" json:"id"`
	Qty       int       `gorm:"not null" json:"qty"`
	Item_name string    `gorm:"not null" json:"item_name"`
	PricingID uuid.UUID `gorm:"not null" json:"pricing_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (i *ItemList) Validate() error {
	validate := validator.New()
	err := validate.Struct(i)
	if err != nil {
		return err
	}

	return nil
}
