/*
 * @Author: fzf404
 * @Date: 2022-01-22 15:24:56
 * @LastEditTime: 2022-02-08 17:09:30
 * @Description: User 管理
 */

package api

import (
	"todo-list/model"
	"todo-list/res"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户注册
 * @param {*gin.Context} c
 */
func Regist(c *gin.Context) {
	var r model.Regist
	// 读取数据
	if err := c.ShouldBind(&r); err != nil {
		res.FailWithDetailed(res.ParamError, err.Error(), "提交信息非法", c)
		return
	}
	service.Regist(r, c)
}

/**
 * @description: 用户登录
 * @param {*gin.Context} c
 */
func Login(c *gin.Context) {
	// 绑定数据
	var l model.Login
	if err := c.ShouldBind(&l); err != nil {
		res.FailWithDetailed(res.ParamError, err.Error(), "用户名或密码错误", c)
		return
	}
	service.Login(l, c)
}
