package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Testimoni struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Image     string    `gorm:"not null" json:"image"`
	Name      string    `gorm:"not null" json:"name" validate:"required,max=255"`
	Event     string    `gorm:"not null" json:"event" validate:"required,max=255"`
	Comment   string    `gorm:"not null" json:"comment" validate:"required"`
	PageID    uint      `json:"page_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (t *Testimoni) Validate() error {
	validate := validator.New()
	err := validate.Struct(t)
	if err != nil {
		return err
	}

	return nil
}
