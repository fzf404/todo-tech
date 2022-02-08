/*
 * @Author: fzf404
 * @Date: 2022-01-22 15:35:53
 * @LastEditTime: 2022-02-08 17:09:51
 * @Description: 鉴权
 */

package middleware

import (
	"todo-list/res"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @description: token 鉴权
 */
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 Token
		tokenString := c.GetHeader("x-token")
		if tokenString == "" {
			res.FailWithMessage(res.NoLogin, "未登录", c)
			c.Abort()
			return
		}

		// 解析 Token
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			res.FailWithMessage(res.TokenExpired, "登录已过期", c)
			c.Abort()
			return
		}

		// 通过验证, 获取 claims 中的 UserID
		uid := claims.UID

		// 安全验证(防止伪造)
		// var u model.User
		// if database.MysqlDB.First(&u, uid).Error != nil {
		// 	res.FailWithMessage(res.TokenExpired, "登录已过期", c)
		// 	c.Abort()
		// 	return
		// }

		// uid 写入上下文
		c.Set("UserID", uid)
		c.Next()
	}
}
