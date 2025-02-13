/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-13 16:44:03
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-13 16:50:40
 * @FilePath: \hospital\middleware\auth.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middleware

import (
	"context"
	"fmt"
	"hospital/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware 验证token的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "未授权访问",
				"error":   "缺少token",
			})
			c.Abort()
			return
		}

		// 从Redis中获取token对应的用户ID
		ctx := context.Background()
		userID, err := utils.Redis.Get(ctx, fmt.Sprintf("token:%s", token)).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "未授权访问",
				"error":   "无效的token",
			})
			c.Abort()
			return
		}

		// 将用户ID存储在上下文中
		c.Set("userID", userID)
		c.Next()
	}
}
