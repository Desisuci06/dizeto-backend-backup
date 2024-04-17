package controller

import (
	model_item_list "dizeto-backend/app/model/item"
	dto "dizeto-backend/app/model/pricing/dto"

	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PricingController struct {
	pricingService service.PricingService
}

func NewPricingController(pricingService service.PricingService) *PricingController {
	return &PricingController{pricingService: pricingService}
}

func (pc *PricingController) CreatePricing(c *gin.Context) {
	pricingID := uuid.New()

	var pricingDTO dto.PricingDTO
	if err := c.ShouldBindJSON(&pricingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create itemList from DTO
	var itemList []*model_item_list.ItemList
	for _, item := range pricingDTO.ItemList {
		newItem := &model_item_list.ItemList{
			ID:        item.ID,
			Qty:       item.Qty,
			Item_name: item.Item_name,
			PricingID: pricingID,
		}
		itemList = append(itemList, newItem)
	}

	// Call service method to create pricing
	if err := pc.pricingService.CreatePricing(pricingID, pricingDTO.Title, pricingDTO.Paket, pricingDTO.Category, itemList, pricingDTO.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (pc *PricingController) GetAllPricing(c *gin.Context) {
	responseDTO, err := pc.pricingService.GetAllPricing()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (pc *PricingController) UpdatePricing(c *gin.Context) {
	id := c.Param("id")

	var pricingDTO dto.PricingDTO
	if err := c.ShouldBindJSON(&pricingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create itemList from DTO
	var itemList []*model_item_list.ItemList
	for _, item := range pricingDTO.ItemList {
		newItem := &model_item_list.ItemList{
			ID:        item.ID,
			Qty:       item.Qty,
			Item_name: item.Item_name,
		}
		itemList = append(itemList, newItem)
	}

	err := pc.pricingService.UpdatePricing(id, itemList, &pricingDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Respond with success message
	utils.SuccessMessage(c, http.StatusOK, "Successfully updated pricing")
}
