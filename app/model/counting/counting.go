package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Counting struct {
	ID                uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	HappyClient       uint      `gorm:"not null" json:"happy_client" validate:"required`
	CompletedProjects uint      `gorm:"not null" json:"completed_projects" validate:"required`
	Subscribers       uint      `gorm:"not null" json:"subscribers" validate:"required`
	PageID            uint      `json:"page_id"`
	CreatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Counting) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	return nil
}
