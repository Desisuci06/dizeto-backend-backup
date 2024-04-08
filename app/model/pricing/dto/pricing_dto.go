package dto

import "github.com/google/uuid"

type PricingDTO struct {
	Title    string `json:"title" binding:"required"`
	Price    uint   `json:"price" binding:"required"`
	Paket    string `json:"paket" binding:"required"`
	Category string `json:"category" binding:"required"`
	ItemList string `json:"item_list" binding:"required"`
}

type ResponseDTO struct {
	ID       uuid.UUID `json:"id"`
	Title    string    `json:"title"`
	Price    uint      `json:"price"`
	Paket    string    `json:"paket"`
	Category string    `json:"category"`
	ItemList string    `json:"item_list"`
}

type ResponsePricingsDTO struct {
	Pricings []*ResponseDTO `json:"pricings"`
}
