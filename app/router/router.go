package router

import (
	"dizeto-backend/app/controller"
	"dizeto-backend/app/repository"
	"dizeto-backend/app/service"
	"dizeto-backend/middleware"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) {
	// Initialize repository
	userRepo := repository.NewUserRepository(db)
	aboutRepo := repository.NewAboutRepository(db)
	highlightRepo := repository.NewHighlightPortofolio(db)
	pricingRepo := repository.NewPricingRepository(db)
	testimoniRepo := repository.NewTestimoniRepository(db)
	countingRepo := repository.NewCountingRepository(db)
	clientRepo := repository.NewClientRepository(db)
	pageRepo := repository.NewLandingRepository(db)
	titleRepo := repository.NewTitleRepository(db)

	// Initialize service
	userService := service.NewUserService(userRepo)
	aboutService := service.NewAboutService(aboutRepo, titleRepo)
	highlightService := service.NewHighlightService(highlightRepo, titleRepo)
	pricingService := service.NewPricingService(pricingRepo, titleRepo)
	testimoniService := service.NewTestimoniService(testimoniRepo)
	countingService := service.NewCountingService(countingRepo)
	clientService := service.NewClientService(clientRepo, titleRepo)
	pageService := service.NewLandingService(pageRepo)
	titleService := service.NewTitleService(titleRepo)

	// Initialize controller
	userController := controller.NewUserController(userService)
	aboutController := controller.NewAboutController(aboutService)
	highlightController := controller.NewHighlightController(highlightService)
	pricingController := controller.NewPricingController(pricingService)
	testimoniController := controller.NewTestimoniController(testimoniService)
	countingController := controller.NewCountingController(countingService)
	clientController := controller.NewClientController(clientService)
	pageController := controller.NewLandingController(pageService)
	titleController := controller.NewTitleController(titleService)

	// Routes
	v1 := r.Group("/api/v1")
	{
		//user
		v1.POST("/register", userController.Register)
		v1.POST("/login", userController.Login)

		//about
		v1.POST("/about", middleware.AuthorizationMiddleware(), aboutController.CreateAbout)
		v1.GET("/about", aboutController.GetAllAbout)
		v1.PUT("/about/:id", middleware.AuthorizationMiddleware(), aboutController.UpdateAbout)

		//highlight
		v1.POST("/highlight_portofolio", middleware.AuthorizationMiddleware(), highlightController.CreateHighlight)
		v1.GET("/highlight_portofolio", highlightController.GetAllHighlight)
		v1.GET("/highlight_portofolio/:id", highlightController.GetHighlightByID)
		v1.PUT("/highlight_portofolio/:id", middleware.AuthorizationMiddleware(), highlightController.UpdateHighlight)

		//pricing
		v1.POST("/pricing", middleware.AuthorizationMiddleware(), pricingController.CreatePricing)
		v1.GET("/pricing", pricingController.GetAllPricing)
		v1.PUT("/pricing/:id", middleware.AuthorizationMiddleware(), pricingController.UpdatePricing)

		//testimoni
		v1.POST("/testimoni", middleware.AuthorizationMiddleware(), testimoniController.CreateTestimoni)
		v1.GET("/testimoni", testimoniController.GetAllTestimoni)
		v1.PUT("/testimoni/:id", middleware.AuthorizationMiddleware(), testimoniController.UpdateTestimoni)

		//counting
		v1.POST("/counting", middleware.AuthorizationMiddleware(), countingController.CreateCounting)
		v1.GET("/counting", countingController.GetAllCounting)
		v1.PUT("/counting/:id", middleware.AuthorizationMiddleware(), countingController.UpdateCounting)

		//client
		v1.POST("/client", middleware.AuthorizationMiddleware(), clientController.CreateClient)
		v1.GET("/client", clientController.GetAllClient)
		v1.PUT("/client/:id", middleware.AuthorizationMiddleware(), clientController.UpdateClient)

		//title
		v1.POST("/title", middleware.AuthorizationMiddleware(), titleController.CreateTitle)
		v1.GET("/title", titleController.GetAllTitle)
		v1.PUT("/title/:id", middleware.AuthorizationMiddleware(), titleController.UpdateTitle)

		//landing page
		v1.GET("/landing-page", pageController.GetLandingPage)
	}
}
