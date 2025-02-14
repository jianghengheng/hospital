/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 13:49:47
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-14 13:36:18
 * @FilePath: \hospital\controllers\user_controller.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package controllers

import (
	"context"
	"fmt"
	"hospital/models"
	"hospital/utils"
	"hospital/utils/response"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"golang.org/x/crypto/bcrypt"
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

// LoginRequest 定义登录请求的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"` // 用户名
	Password string `json:"password" binding:"required" example:"123456"`  // 密码
}

type UserController struct{}

// Create godoc
// @Summary 创建新用户
// @Description 创建一个新的用户账号
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param user body models.User true "用户信息"
// @Success 200 {object} UserResponse "成功创建用户"
// @Failure 400 {object} Response "请求参数错误"
// @Failure 401 {object} Response "未授权访问"
// @Failure 500 {object} Response "服务器内部错误"
// @Router /users [post]
func (u *UserController) Create(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.BadRequest(c, "请求参数错误", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.ServerError(c, "密码加密失败", err.Error())
		return
	}
	user.Password = string(hashedPassword)

	if err := utils.DB.Create(&user).Error; err != nil {
		response.ServerError(c, "创建用户失败", err.Error())
		return
	}

	response.Success(c, user, "用户创建成功")
}

// Get godoc
// @Summary 获取用户信息
// @Description 根据用户ID获取用户详细信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "用户ID"
// @Success 200 {object} UserResponse "成功获取用户信息"
// @Failure 401 {object} Response "未授权访问"
// @Failure 404 {object} Response "用户不存在"
// @Router /users/{id} [get]
func (u *UserController) Get(c *gin.Context) {
	id := c.Param("id")
	var user models.User

	if err := utils.DB.Where("id = ?", id).First(&user).Error; err != nil {
		response.NotFound(c, "用户不存在", err.Error())
		return
	}

	response.Success(c, user, "获取用户成功")
}
func a() {

	var mutex sync.Mutex
	counter := 0
	workers := 5
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录接口
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param request body LoginRequest true "登录信息"
// @Success 200 {object} Response
// @Router /login [post]
func (u *UserController) Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		response.BadRequest(c, "参数错误", err.Error())
		return
	}

	var user models.User
	if err := utils.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		response.Unauthorized(c, "用户名或密码错误", "")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		response.Unauthorized(c, "用户名或密码错误", "")
		return
	}

	token := uuid.New().String()
	ctx := context.Background()

	if err := utils.Redis.Set(ctx, fmt.Sprintf("token:%s", token), user.ID, 24*time.Hour).Err(); err != nil {
		response.ServerError(c, "登录失败", err.Error())
		return
	}

	response.Success(c, gin.H{
		"token": token,
		"user":  user,
	}, "登录成功")
}

// Export 导出用户数据为Excel
// @Summary 导出用户数据为Excel
// @Description 导出所有用户数据为Excel文件
// @Tags 用户管理
// @Accept json
// @Produce application/octet-stream
// @Security Bearer
// @Success 200 {file} file "成功导出Excel文件"
// @Failure 401 {object} Response "未授权访问"
// @Failure 500 {object} Response "服务器内部错误"
// @Router /export [get]
func (u *UserController) Export(c *gin.Context) {
	var users []models.User
	if err := utils.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error(), Status: http.StatusInternalServerError, Message: "服务器内部错误"})
		return
	}

	f := excelize.NewFile()
	// 创建一个工作表
	index, err := f.NewSheet("Users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error(), Status: http.StatusInternalServerError, Message: "服务器内部错误"})
		return
	}

	// 设置表头
	headers := []string{"ID", "Username", "Email", "HeadImage", "Phone"}
	for i, header := range headers {
		cell, err := excelize.CoordinatesToCellName(1+i, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Response{Error: err.Error(), Status: http.StatusInternalServerError, Message: "服务器内部错误"})
			return
		}
		f.SetCellValue("Users", cell, header)
	}

	// 填充数据
	for i, user := range users {
		row := i + 2 // 从第二行开始填充数据
		f.SetCellValue("Users", fmt.Sprintf("A%d", row), user.ID)
		f.SetCellValue("Users", fmt.Sprintf("B%d", row), user.Username)
		f.SetCellValue("Users", fmt.Sprintf("C%d", row), user.Email)
		f.SetCellValue("Users", fmt.Sprintf("D%d", row), user.HeadImage)
		f.SetCellValue("Users", fmt.Sprintf("E%d", row), user.Phone)
	}

	// 设置默认工作表
	f.SetActiveSheet(index)

	// 写入文件
	filename := "users.xlsx"
	if err := f.SaveAs(filename); err != nil {
		c.JSON(http.StatusInternalServerError, Response{Error: err.Error(), Status: http.StatusInternalServerError, Message: "服务器内部错误"})
		return
	}

	// 发送文件给客户端
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.File(filename)

	// 删除临时文件
	os.Remove(filename)
}
