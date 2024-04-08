package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type About struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title       string    `gorm:"not null" json:"title" validate:"required,min=3,max=255"`
	Subtitle    string    `gorm:"not null" json:"subtitle" validate:"required,min=3,max=255"`
	Description string    `gorm:"not null" json:"description" validate:"required"`
	Note        string    `gorm:"not null" json:"note"`
	Image       string    `gorm:"not null" json:"image"`
	PageID      uint      `json:"page_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *About) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		return err
	}

	return nil
}
