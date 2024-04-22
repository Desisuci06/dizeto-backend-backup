package dto

import "github.com/google/uuid"

type HighlightDTO struct {
	// Title string `json:"title" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type ResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Image string    `json:"image"`
}

type ResponseHighlightsDTO struct {
	Highlight_portofolio []*ResponseDTO `json:"highlight_portofolio"`
}
