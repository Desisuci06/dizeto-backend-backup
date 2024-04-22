package dto

type TitleDTO struct {
	ID      int    `json:"id"`
	KdTitle string `json:"kd_title" binding:"required"`
	NmTitle string `json:"nm_title" binding:"required"`
}

type ResponseDTO struct {
	ID      int    `json:"id"`
	KdTitle string `json:"kd_title"`
	NmTitle string `json:"nm_title"`
}

type ResponseTitlesDTO struct {
	Titles []*ResponseDTO `json:"Titles"`
}
