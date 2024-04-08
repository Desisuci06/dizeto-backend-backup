package controller

import (
	dto "dizeto-backend/app/model/counting/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CountingController struct {
	countingService service.CountingService
}

func NewCountingController(countingService service.CountingService) *CountingController {
	return &CountingController{countingService: countingService}
}

func (cc *CountingController) CreateCounting(c *gin.Context) {
	var countingDTO dto.CountingDTO
	if err := c.ShouldBindJSON(&countingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.countingService.CreateCounting(countingDTO.HappyClient, countingDTO.CompletedProjects, countingDTO.Subscribers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (cc *CountingController) GetAllCounting(c *gin.Context) {
	responseDTO, err := cc.countingService.GetAllCounting()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (cc *CountingController) UpdateCounting(c *gin.Context) {
	id := c.Param("id")

	var countingDTO dto.CountingDTO
	if err := c.ShouldBindJSON(&countingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.countingService.UpdateCounting(id, countingDTO.HappyClient, countingDTO.CompletedProjects, countingDTO.Subscribers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
