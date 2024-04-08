package controller

import (
	"dizeto-backend/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LandingController struct {
	landingService service.LandingService
}

func NewLandingController(landingService service.LandingService) *LandingController {
	return &LandingController{landingService: landingService}
}

func (lc *LandingController) GetLandingPage(c *gin.Context) {
	pages, err := lc.landingService.GetLandingPage()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pages)
}
