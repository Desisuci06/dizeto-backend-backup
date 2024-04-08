package dto

import "github.com/google/uuid"

type AboutDTO struct {
	Title       string `json:"title" binding:"required"`
	Subtitle    string `json:"subtitle" binding:"required"`
	Description string `json:"description" binding:"required"`
	Note        string `json:"note" binding:"required"`
	Image       string `json:"image" binding:"required"`
}

type ResponseDTO struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Subtitle    string    `json:"subtitle"`
	Description string    `json:"description"`
	Note        string    `json:"note"`
	Image       string    `json:"image"`
}

type ResponseAboutsDTO struct {
	Abouts []*ResponseDTO `json:"abouts"`
}
