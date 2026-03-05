package routes

import (
	"amar_dokan/controllers"
	"amar_dokan/middleware"
	"amar_dokan/repositories"

	"github.com/gin-gonic/gin"
)

func registerCategory(router *gin.RouterGroup, controller controllers.CategoryController, userRepo repositories.UserRepository, jwtSecret string) {
	grp := router.Group("/category")
	grp.Use(middleware.AuthRequired(jwtSecret, userRepo))

	grp.POST("/", controller.Create)
	// grp.GET("/", controller.Get)
	// grp.PUT("/:id", controller.Update)
	// grp.DELETE("/:id", controller.Delete)
	//	productGr := rg.Group("/products")

}
