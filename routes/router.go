package routes

import (
	"amar_dokan/config"
	"amar_dokan/controllers"
	productcontroller "amar_dokan/controllers/product_controller"
	usercontroller "amar_dokan/controllers/user_controller"
	"amar_dokan/middleware"
	"amar_dokan/repositories"
	"amar_dokan/services"
	productservice "amar_dokan/services/product_service"
	userService "amar_dokan/services/user_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	addToCartRepo := repositories.NewAddToCartRepository(db)
	caregoryRepo := repositories.NewCategoryRepo(db)

	userSvc := userService.NewUserService(userRepo, cfg)
	productSvc := productservice.NewProductService(cfg.JwtSecureKey, productRepo)
	addToCardSvc := services.NewAddToCardService(cfg.JwtSecureKey, addToCartRepo)
	categorySvc := services.NewCategoryService(&caregoryRepo, cfg.JwtSecureKey)

	userCtrl := usercontroller.NewUserController(userSvc)
	productCtrl := productcontroller.NewProductController(productSvc)
	addToCardCtrl := controllers.NewAddToCardController(addToCardSvc)
	categoryCtrl := controllers.NewCategoryController(categorySvc)

	r := gin.New()

	// Global middleware
	r.Use(middleware.CORS())
	r.Use(middleware.RateLimiter())
	r.Use(gin.Recovery())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": cfg.ServiceName,
			"version": cfg.Version,
		})
	})

	// ── API v1 ───────────────────────────────────────────────────────────
	apiV1 := r.Group("/api/v1")

	registerUserRoutes(apiV1, userCtrl, userRepo, cfg.JwtSecureKey)
	registerProductRoutes(apiV1, productCtrl, userRepo, cfg.JwtSecureKey)
	registerAddToCartRoutes(apiV1, addToCardCtrl, userRepo, cfg.JwtSecureKey)
	registerCategory(apiV1, categoryCtrl, userRepo, cfg.JwtSecureKey)

	return r
}
