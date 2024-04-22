package controller

import (
	dto "dizeto-backend/app/model/title/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TitleController struct {
	titleService service.TitleService
}

func NewTitleController(titleService service.TitleService) *TitleController {
	return &TitleController{titleService: titleService}
}

func (ac *TitleController) CreateTitle(c *gin.Context) {
	var titleDTO dto.TitleDTO
	if err := c.ShouldBindJSON(&titleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.titleService.CreateTitle(titleDTO.ID, titleDTO.KdTitle, titleDTO.NmTitle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (ac *TitleController) GetAllTitle(c *gin.Context) {
	responseDTO, err := ac.titleService.GetAllTitle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (ac *TitleController) UpdateTitle(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var titleDTO dto.TitleDTO
	if err := c.ShouldBindJSON(&titleDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.titleService.UpdateTitle(idInt, titleDTO.KdTitle, titleDTO.NmTitle); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
