package controller

import (
	dto "dizeto-backend/app/model/pricing/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PricingController struct {
	pricingService service.PricingService
}

func NewPricingController(pricingService service.PricingService) *PricingController {
	return &PricingController{pricingService: pricingService}
}

func (pc *PricingController) CreatePricing(c *gin.Context) {
	var pricingDTO dto.PricingDTO
	if err := c.ShouldBindJSON(&pricingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.pricingService.CreatePricing(pricingDTO.Title, pricingDTO.Paket, pricingDTO.Category, pricingDTO.ItemList, pricingDTO.Price); err != nil {
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

	if err := pc.pricingService.UpdatePricing(id, pricingDTO.Title, pricingDTO.Paket, pricingDTO.Category, pricingDTO.ItemList, pricingDTO.Price); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
