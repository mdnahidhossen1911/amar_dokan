package routes

import (
	"amar_dokan/config"
	productcontroller "amar_dokan/controllers/product_controller"
	usercontroller "amar_dokan/controllers/user_controller"
	"amar_dokan/middleware"
	"amar_dokan/repositories"
	productservice "amar_dokan/services/product_service"
	userService "amar_dokan/services/user_service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(cfg *config.Config, db *gorm.DB) *gin.Engine {

	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)

	userSvc := userService.NewUserService(userRepo, cfg)
	productSvc := productservice.NewProductService(cfg.JwtSecureKey, productRepo)

	userCtrl := usercontroller.NewUserController(userSvc)
	productCtrl := productcontroller.NewProductController(productSvc)

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
	api := r.Group("/api/v1")

	registerUserRoutes(api, userCtrl, userRepo, cfg.JwtSecureKey)
	registerProductRoutes(api, productCtrl, userRepo, cfg.JwtSecureKey)
	return r
}
