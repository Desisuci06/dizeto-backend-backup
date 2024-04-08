package dto

import "github.com/google/uuid"

type CountingDTO struct {
	HappyClient       uint `json:"happy_client" binding:"required"`
	CompletedProjects uint `json:"completed_projects" binding:"required"`
	Subscribers       uint `json:"subscribers" binding:"required"`
}

type ResponseDTO struct {
	ID                uuid.UUID `json:"id"`
	HappyClient       uint      `json:"happy_client" `
	CompletedProjects uint      `json:"completed_projects" `
	Subscribers       uint      `json:"subscribers" `
}

type ResponseCountingsDTO struct {
	Countings []*ResponseDTO `json:"counting"`
}
