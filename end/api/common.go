/*
 * @Author: fzf404
 * @Date: 2022-01-22 15:23:36
 * @LastEditTime: 2022-02-08 17:09:35
 * @Description: 测试
 */
package api

import (
	"todo-list/res"

	"github.com/gin-gonic/gin"
)

/**
 * @description: Ping 接口
 * @param {*gin.Context} c
 */
func Ping(c *gin.Context) {
	res.OkWithMessage("pong", c)
}

/**
 * @description: 登录验证
 * @param {*gin.Context} c
 */
func LoginVerify(c *gin.Context) {
	res.OkWithMessage("已登录", c)
}
