/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:28:51
 * @LastEditTime: 2022-02-08 17:10:55
 * @Description: User
 */

package model

import "gorm.io/gorm"

type Regist struct {
	Username string `json:"username" binding:"min=3,max=20,required"` // 登录名
	Password string `json:"password" binding:"min=6,max=20,required"` // 密码
	Email    string `json:"email" binding:"email,required"`           // 邮箱
}

type Login struct {
	Username string `json:"username" binding:"min=3,max=20,required"`
	Password string `json:"password" binding:"min=6,max=20,required"`
}

/**
 * @description: users 数据表
 */
type User struct {
	gorm.Model
	Username string // 登录名
	Password string // 密码
	Email    string // 邮箱
	Avatar   string `gorm:"default:'https://www.fzf404.top/public/avatar.jpg'"` // 头像
	Status   bool   `gorm:"default:false"`                                      // 是否激活
}

type UserDto struct {
	ID       uint   `json:"uid"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Status   bool   `json:"status"`
}
