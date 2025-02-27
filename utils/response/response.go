/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-14 09:58:21
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-14 10:10:29
 * @FilePath: \hospital\utils\response\response.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 通用响应结构
// @Description API的通用响应格式
type Response struct {
	// 响应数据
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Status  int         `json:"status,omitempty"`
	// 错误信息
	Error string `json:"error,omitempty"`
}

// Success 成功响应
func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, status int, message string, err string) {
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Error:   err,
	})
}

// Unauthorized 未授权响应
func Unauthorized(c *gin.Context, message string, err string) {
	Error(c, http.StatusUnauthorized, message, err)
}

// BadRequest 请求参数错误响应
func BadRequest(c *gin.Context, message string, err string) {
	Error(c, http.StatusBadRequest, message, err)
}

// NotFound 资源不存在响应
func NotFound(c *gin.Context, message string, err string) {
	Error(c, http.StatusNotFound, message, err)
}

// ServerError 服务器内部错误响应
func ServerError(c *gin.Context, message string, err string) {
	Error(c, http.StatusInternalServerError, message, err)
}
