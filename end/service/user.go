/*
 * @Author: fzf404
 * @Date: 2022-01-22 22:00:12
 * @LastEditTime: 2022-02-08 17:11:06
 * @Description: User 处理
 */
package service

import (
	"log"
	"todo-list/database"
	"todo-list/model"
	"todo-list/res"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @description: 用户登录
 */
func Login(l model.Login, c *gin.Context) {

	// 验证用户存在
	var u model.User
	if err := database.MysqlDB.Where("username = ?", l.Username).First(&u).Error; err != nil {
		if err := database.MysqlDB.Where("email = ?", l.Username).First(&u).Error; err != nil {
			res.FailWithMessage(res.ParamError, "用户名或密码错误", c)
			return
		}
	}

	// 密码加密存储
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		res.FailWithMessage(res.ParamError, "用户名或密码错误", c)
		return
	}

	// 生成 Token
	token, err := utils.ReleaseToken(u)
	if err != nil {
		log.Print("Token generate error:", err)
		res.FailWithMessage(res.TokenErr, "Token分发错误", c)
		return
	}

	// 数据转换
	var ud model.UserDto
	utils.StructAssign(&ud, &u)
	ud.ID = u.ID

	// 成功响应
	res.OkWithDetailed(gin.H{"user": ud, "token": token}, "登录成功", c)
}

/**
 * @description: 用户注册
 */

func Regist(r model.Regist, c *gin.Context) {

	// 判断 用户名 & 邮箱 是否存在
	var count int64
	database.MysqlDB.Model(&model.User{}).Where("username = ?", r.Username).Count(&count)
	if count != 0 {
		res.FailWithMessage(res.ParamError, "用户名已存在", c)
		return
	}
	database.MysqlDB.Model(&model.User{}).Where("email = ?", r.Email).Count(&count)
	if count != 0 {
		res.FailWithMessage(res.ParamError, "邮箱已存在", c)
		return
	}

	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password hash error:", err)
		res.FailWithMessage(res.HashErr, "密码加密失败", c)
		return
	}
	r.Password = string(hashPassword)

	// 结构体转换
	var u model.User
	utils.StructAssign(&u, &r)

	// 插入数据
	if database.MysqlDB.Create(&u).Error != nil {
		res.FailWithMessage(res.ServerErr, "数据库连接失败", c)
		return
	}

	res.OkWithMessage("注册成功", c)
}
