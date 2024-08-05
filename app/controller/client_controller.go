package controller

import (
	dto "dizeto-backend/app/model/client/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	clientService service.ClientService
}

func NewClientController(clientService service.ClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

func (hc *ClientController) CreateClient(c *gin.Context) {
	var clientDTO dto.ClientDTO
	if err := c.ShouldBindJSON(&clientDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hc.clientService.CreateClient("CLIE", clientDTO.LogoURL, clientDTO.Theme, clientDTO.Href, clientDTO.Alt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Successfully")

}

func (hc *ClientController) GetAllClient(c *gin.Context) {
	responseDTO, err := hc.clientService.GetAllClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, responseDTO)
}

func (hc *ClientController) GetClientByID(c *gin.Context) {
	id := c.Param("id")

	client, err := hc.clientService.GetClientByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessData(c, http.StatusOK, client)
}

func (hc *ClientController) UpdateClient(c *gin.Context) {
	id := c.Param("id")

	var clientDTO dto.ClientDTO
	if err := c.ShouldBindJSON(&clientDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := hc.clientService.UpdateClient(id, "CLIE", clientDTO.LogoURL, clientDTO.Theme, clientDTO.Href, clientDTO.Alt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Mengembalikan respons berhasil
	utils.SuccessMessage(c, http.StatusOK, "Successfully")
}
