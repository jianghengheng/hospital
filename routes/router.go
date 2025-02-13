/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 13:49:49
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-13 17:45:43
 * @FilePath: \hospital\routes\router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routes

import (
	"hospital/controllers"
	"hospital/middleware"

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

		// 登录路由 - 不需要验证
		api.POST("/login", userController.Login)
		api.GET("/export", userController.Export) // 添加导出路由

		// 需要验证的路由组
		auth := api.Group("")
		auth.Use(middleware.AuthMiddleware())
		{
			// 用户组路由
			users := auth.Group("/users")
			{
				users.POST("/", userController.Create)
				users.GET("/:id", userController.Get)

			}
		}
	}

	return r
}
