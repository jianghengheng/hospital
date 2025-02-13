/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 13:49:51
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-13 17:01:58
 * @FilePath: \hospital\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"hospital/config"
	"hospital/docs" // 导入生成的docs
	"hospital/routes"
	"hospital/utils"
	"log"
	"time"
)

// @title Hospital API
// @version 1.0
// @description This is a hospital management system server.
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// 添加启动时间日志
	startTime := time.Now()
	log.Printf("服务启动时间: %v", startTime.Format("2006-01-02 15:04:05"))

	// 初始化swagger文档
	docs.SwaggerInfo.Title = "Hospital API"
	docs.SwaggerInfo.Description = "This is a hospital management system server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		log.Fatalf("配置初始化失败: %v", err)
	}
	log.Println("配置初始化成功")

	// 初始化MySQL
	if err := utils.InitMySQL(); err != nil {
		log.Fatalf("MySQL初始化失败: %v", err)
	}
	log.Println("MySQL初始化成功")

	// 初始化Redis
	if err := utils.InitRedis(); err != nil {
		log.Fatalf("Redis初始化失败: %v", err)
	}
	log.Println("Redis初始化成功")

	// 设置路由
	r := routes.SetupRouter()
	log.Println("路由设置成功")

	// 打印服务启动信息
	fmt.Println("\n=================================")
	fmt.Println("  Hospital API 服务已启动")
	fmt.Println("  Swagger文档: http://localhost:8080/swagger/index.html")
	fmt.Printf("  启动用时: %v\n", time.Since(startTime))
	fmt.Println("=================================\n")

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
