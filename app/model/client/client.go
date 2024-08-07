package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Title     string    `gorm:"not null" json:"title" validate:"required,min=3,max=255"`
	LogoURL   string    `gorm:"not null" json:"logoURL"`
	Theme     string    `gorm:"not null" json:"theme"`
	Href      string    `gorm:"not null" json:"href"`
	Alt       string    `gorm:"not null" json:"alt"`
	PageID    uint      `json:"page_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Client) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
