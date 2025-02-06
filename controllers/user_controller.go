/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 13:49:47
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-06 17:05:37
 * @FilePath: \hospital\controllers\user_controller.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controllers

import (
	"hospital/models"
	"hospital/utils"
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

// UserResponse 用户响应结构
// @Description 包含用户数据的响应
type UserResponse struct {
	// 响应数据
	Data models.User `json:"data"`
}

type UserController struct{}

// Create godoc
// @Summary 创建新用户
// @Description 创建一个新的用户账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} UserResponse "成功创建用户"
// @Failure 400 {object} Response "请求参数错误"
// @Failure 500 {object} Response "服务器内部错误"
// @Router /users [post]
func (u *UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{Error: err.Error(), Status: http.StatusBadRequest, Message: "请求参数错误"})
		return
	}

	if err := utils.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error(), Status: http.StatusInternalServerError, Message: "服务器内部错误"})
		return
	}

	c.JSON(http.StatusOK, Response{Data: user, Status: http.StatusOK, Message: "success"})
}

// Get godoc
// @Summary 获取用户信息
// @Description 根据用户ID获取用户详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path string true "用户ID"
// @Success 200 {object} UserResponse "成功获取用户信息"
// @Failure 404 {object} Response "用户不存在"
// @Router /users/{id} [get]
func (u *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := utils.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{Error: "用户不存在"})
		return
	}

	c.JSON(http.StatusOK, Response{Data: user})
}
