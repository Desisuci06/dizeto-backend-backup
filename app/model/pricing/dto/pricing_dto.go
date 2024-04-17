package dto

import (
	model_item_list "dizeto-backend/app/model/item"
	item_list_dto "dizeto-backend/app/model/item/dto"

	"github.com/google/uuid"
)

type PricingDTO struct {
	Title    string                   `json:"title" binding:"required"`
	Price    uint                     `json:"price" binding:"required"`
	Paket    string                   `json:"paket" binding:"required"`
	Category string                   `json:"category" binding:"required"`
	ItemList []*item_list_dto.ItemDTO `json:"item_list" binding:"required"`
}

type ResponseDTO struct {
	ID       uuid.UUID                   `json:"id"`
	Title    string                      `json:"title"`
	Price    uint                        `json:"price"`
	Paket    string                      `json:"paket"`
	Category string                      `json:"category"`
	ItemList []*model_item_list.ItemList `json:"item_list"`
}

type ResponsePricingsDTO struct {
	Pricings []*ResponseDTO `json:"pricings"`
}
