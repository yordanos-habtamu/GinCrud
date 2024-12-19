package routes

import (
	"github.com/yordanos-habtamu/GinCrud/controller"
	"github.com/yordanos-habtamu/GinCrud/middleware"
	"github.com/gin-gonic/gin"
)

func SetupBookRoutes(r *gin.Engine) {
	bookRoutes := r.Group("/books")
	{
		bookRoutes.GET("/", controllers.GetBooks)
		bookRoutes.GET("/:id", controllers.GetBook)
		bookRoutes.POST("/", controllers.CreateBook)
		bookRoutes.PUT("/:id", controllers.UpdateBook)
		bookRoutes.DELETE("/:id", controllers.DeleteBook)
	}
	bookRoutes.Use(middlewares.AuthMiddleware()) // Example middleware usage
}
