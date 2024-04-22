package controller

import (
	dto "dizeto-backend/app/model/about/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AboutController struct {
	aboutService service.AboutService
}

func NewAboutController(aboutService service.AboutService) *AboutController {
	return &AboutController{aboutService: aboutService}
}

func (ac *AboutController) CreateAbout(c *gin.Context) {
	var aboutDTO dto.AboutDTO
	if err := c.ShouldBindJSON(&aboutDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.aboutService.CreateAbout("ABT", aboutDTO.Subtitle, aboutDTO.Description, aboutDTO.Note, aboutDTO.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (ac *AboutController) GetAllAbout(c *gin.Context) {
	responseDTO, err := ac.aboutService.GetAllAbout()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (ac *AboutController) UpdateAbout(c *gin.Context) {
	id := c.Param("id")

	var aboutDTO dto.AboutDTO
	if err := c.ShouldBindJSON(&aboutDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ac.aboutService.UpdateAbout(id, "ABT", aboutDTO.Subtitle, aboutDTO.Description, aboutDTO.Note, aboutDTO.Image); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
