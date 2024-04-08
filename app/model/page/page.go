package model

import (
	model_about "dizeto-backend/app/model/about"
	model_client "dizeto-backend/app/model/client"
	model_counting "dizeto-backend/app/model/counting"
	model_highlight "dizeto-backend/app/model/highlight_porto"
	model_pricing "dizeto-backend/app/model/pricing"
	model_testimoni "dizeto-backend/app/model/testimoni"
	"time"
)

type Page struct {
	ID         uint                                  `gorm:"primary_key" json:"id"`
	Title      string                                `gorm:"not null" json:"title"`
	Abouts     []model_about.About                   `gorm:"foreignKey:PageID" json:"abouts"`
	Clients    []model_client.Client                 `gorm:"foreignKey:PageID" json:"clients"`
	Countings  []model_counting.Counting             `gorm:"foreignKey:PageID" json:"countings"`
	Highlights []model_highlight.HighlightPortofolio `gorm:"foreignKey:PageID" json:"highlights"`
	Pricings   []model_pricing.Pricing               `gorm:"foreignKey:PageID" json:"pricings"`
	Testimonis []model_testimoni.Testimoni           `gorm:"foreignKey:PageID" json:"testimonis"`
	CreatedAt  time.Time                             `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time                             `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
