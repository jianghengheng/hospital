/*
 * @Author: jiangheng jh@pzds.com
 * @Date: 2025-02-06 13:49:44
 * @LastEditors: jiangheng jh@pzds.com
 * @LastEditTime: 2025-02-06 16:21:10
 * @FilePath: \hospital\models\user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package models

import (
	"hospital/utils"
)

// User 用户模型
// @Description 用户信息结构体
type User struct {
	// 用户ID
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	// 用户名
	// @Description 用户的登录名，必须唯一
	Username string `json:"username" gorm:"type:varchar(100);unique;not null" example:"johndoe" binding:"required"`
	// 密码
	// @Description 用户的登录密码，不能为空
	Password string `json:"password" gorm:"type:varchar(100);not null" example:"password123" binding:"required"`
	// 邮箱
	// @Description 用户的邮箱地址，必须唯一
	Email string `json:"email" gorm:"type:varchar(100);unique" example:"john@example.com" binding:"required,email"`
}

func init() {
	utils.RegisterModel(&User{})
}
