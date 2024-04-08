package dto

import "github.com/google/uuid"

type TestimoniDTO struct {
	Image   string `json:"image" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Event   string `json:"event" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

type ResponseDTO struct {
	ID      uuid.UUID `json:"id"`
	Image   string    `json:"image"`
	Name    string    `json:"name"`
	Event   string    `json:"event"`
	Comment string    `json:"comment"`
}

type ResponseTestimonisDTO struct {
	Testimonis []*ResponseDTO `json:"testimoni"`
}
