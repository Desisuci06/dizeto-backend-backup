package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type HighlightPortofolio struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title     string    `gorm:"not null" json:"title" validate:"required,min=3,max=255"`
	Image     string    `gorm:"not null" json:"image"`
	PageID    uint      `json:"page_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (a *HighlightPortofolio) Validate() error {
	validate := validator.New()
	err := validate.Struct(a)
	if err != nil {
		return err
	}

	return nil
}
