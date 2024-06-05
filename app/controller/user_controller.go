package controller

import (
	"dizeto-backend/app/model/user/dto"
	"dizeto-backend/app/service"
	"dizeto-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{userService}
}

func (uc *UserController) Register(c *gin.Context) {
	// Parse request body
	var req dto.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to register user
	err := uc.userService.Register(req.Username, req.Password, req.FirstName, req.LastName, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	utils.SuccessMessage(c, http.StatusOK, "User registered Successfully")
}

func (uc *UserController) Login(c *gin.Context) {
	// Parse request body
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call service to authenticate user
	user, token, err := uc.userService.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		// id, username, first_name, last_name, email, role
		"Id":       user.ID,
		"Username": user.FirstName,
		"Lastname": user.LastName,
		"email":    user.Email,
		"role":     user.Role,
		"token":    token,
	})
}
