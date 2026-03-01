package routes

import (
	productcontroller "amar_dokan/controllers/product_controller"
	"amar_dokan/middleware"
	"amar_dokan/repositories"

	"github.com/gin-gonic/gin"
)

func registerProductRoutes(rg *gin.RouterGroup, ctrl productcontroller.ProductController, userRepo repositories.UserRepository, jwtSecret string) {
	productGr := rg.Group("/products")
	productGr.GET("", ctrl.Get)

	product := productGr.Group("")
	product.Use(middleware.AuthRequired(jwtSecret, userRepo))
	{
		product.POST("", ctrl.Create)
		product.PUT("/:id", ctrl.Update)
		product.DELETE("/:id", ctrl.Delete)
	}
}
