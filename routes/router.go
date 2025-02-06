package routes

import (
	"hospital/controllers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger 文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API 路由组
	api := r.Group("/api")
	{
		// 用户相关路由
		userController := controllers.UserController{}
		users := api.Group("/users")
		{
			users.POST("/", userController.Create)
			users.GET("/:id", userController.Get)
		}
	}

	return r
}
