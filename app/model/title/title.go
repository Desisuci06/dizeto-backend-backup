package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Title struct {
	ID        int       `gorm:"primary_key" json:"id"`
	KdTitle   string    `gorm:"not null" json:"kd_title" validate:"required"`
	NmTitle   string    `gorm:"not null" json:"nm_title" validate:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Title) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return err
	}

	return nil
}
