package routes

import (
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/controllers"
	"github.com/LuanFreitasRibeiro/webapi-mvc-golang.git/server/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/", controllers.GetUsers)
		}

		main.POST("login", controllers.Login)

		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowAllBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
