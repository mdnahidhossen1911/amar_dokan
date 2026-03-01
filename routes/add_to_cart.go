package routes

import (
	"amar_dokan/controllers"
	"amar_dokan/middleware"
	"amar_dokan/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterAddToCartRoutes(router *gin.RouterGroup, controller controllers.AddToCardController, userRepo repositories.UserRepository, jwtSecret string) {
	grp := router.Group("/add-to-cart")
	grp.Use(middleware.AuthRequired(jwtSecret, userRepo))

	grp.POST("/", controller.Create)
	grp.GET("/", controller.Get)
	grp.PUT("/", controller.Update)
	grp.DELETE("/", controller.Delete)

}
