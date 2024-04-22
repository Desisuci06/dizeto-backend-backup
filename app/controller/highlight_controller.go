package controller

import (
	dto "dizeto-backend/app/model/highlight_porto/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HighlightController struct {
	highlightService service.HighlightService
}

func NewHighlightController(highlightService service.HighlightService) *HighlightController {
	return &HighlightController{highlightService: highlightService}
}

func (hc *HighlightController) CreateHighlight(c *gin.Context) {
	var highlightDTO dto.HighlightDTO
	if err := c.ShouldBindJSON(&highlightDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hc.highlightService.CreateHighlight("PORT", highlightDTO.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (hc *HighlightController) GetAllHighlight(c *gin.Context) {
	responseDTO, err := hc.highlightService.GetAllHighlight()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (hc *HighlightController) GetHighlightByID(c *gin.Context) {
	id := c.Param("id")

	highlight, err := hc.highlightService.GetHighlightByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, highlight)
}

func (hc *HighlightController) UpdateHighlight(c *gin.Context) {
	id := c.Param("id")

	var highlightDTO dto.HighlightDTO
	if err := c.ShouldBindJSON(&highlightDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hc.highlightService.UpdateHighlight(id, "PORT", highlightDTO.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
