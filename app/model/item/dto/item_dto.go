package dto

type ItemDTO struct {
	ID        int    `json:"id" binding:"required"`
	Qty       int    `json:"qty" binding:"required"`
	Item_name string `json:"item_name" binding:"required"`
}

// type ResponseDTO struct {
// 	ID    uuid.UUID `json:"id"`
// 	Title string    `json:"title"`
// 	Image string    `json:"image"`
// }

// type ResponseHighlightsDTO struct {
// 	Highlight_portofolio []*ResponseDTO `json:"highlight_portofolio"`
// }