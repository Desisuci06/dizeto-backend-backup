package dto

import "github.com/google/uuid"

type ClientDTO struct {
	Title string `json:"title" binding:"required"`
	Image string `json:"image" binding:"required"`
}

type ResponseDTO struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
	Image string    `json:"image"`
}

type ResponseClientsDTO struct {
	Clients []*ResponseDTO `json:"client"`
}
