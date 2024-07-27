package dto

import "github.com/google/uuid"

type ItemDTO struct {
	ID        uuid.UUID `json:"id"`
	Qty       int       `json:"qty" binding:"required"`
	Item_name string    `json:"item_name" binding:"required"`
}

// type ResponseDTO struct {
// 	ID    uuid.UUID `json:"id"`
// 	Title string    `json:"title"`
// 	Image string    `json:"image"`
// }

// type ResponseHighlightsDTO struct {
// 	Highlight_portofolio []*ResponseDTO `json:"highlight_portofolio"`
// }
