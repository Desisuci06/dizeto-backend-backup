package controller

import (
	dto "dizeto-backend/app/model/testimoni/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TestimoniController struct {
	testimoniService service.TestimoniService
}

func NewTestimoniController(testimoniService service.TestimoniService) *TestimoniController {
	return &TestimoniController{testimoniService: testimoniService}
}

func (pc *TestimoniController) CreateTestimoni(c *gin.Context) {
	var testimoniDTO dto.TestimoniDTO
	if err := c.ShouldBindJSON(&testimoniDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.testimoniService.CreateTestimoni(testimoniDTO.Image, testimoniDTO.Name, testimoniDTO.Event, testimoniDTO.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (pc *TestimoniController) GetAllTestimoni(c *gin.Context) {
	responseDTO, err := pc.testimoniService.GetAllTestimoni()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (pc *TestimoniController) UpdateTestimoni(c *gin.Context) {
	id := c.Param("id")

	var testimoniDTO dto.TestimoniDTO
	if err := c.ShouldBindJSON(&testimoniDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.testimoniService.UpdateTestimoni(id, testimoniDTO.Image, testimoniDTO.Name, testimoniDTO.Event, testimoniDTO.Comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
