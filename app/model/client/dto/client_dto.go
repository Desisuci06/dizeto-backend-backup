package dto

import "github.com/google/uuid"

type ClientDTO struct {
	// Title string `json:"title" binding:"required"`
	Title   string `json:"title"`
	LogoURL string `json:"logoURL"`
	Theme   string `json:"theme"`
	Href    string `json:"href"`
	Alt     string `json:"alt"`
}

type ResponseDTO struct {
	ID      uuid.UUID `json:"id"`
	Title   string    `json:"title"`
	LogoURL string    `json:"logoURL"`
	Theme   string    `json:"theme"`
	Href    string    `json:"href"`
	Alt     string    `json:"alt"`
}

type ResponseClientsDTO struct {
	Clients []*ResponseDTO `json:"client"`
}
