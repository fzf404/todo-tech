/*
 * @Author: fzf404
 * @Date: 2021-09-22 14:16:48
 * @LastEditTime: 2022-02-08 17:09:46
 * @Description: 跨域
 */
package middleware

import (
	"regexp"
	"todo-list/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/**
* @description: 跨域
 */
func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	c.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie", "x-token"}
	if gin.Mode() == gin.ReleaseMode {
		// 生产环境配置跨域域名，否则 403
		c.AllowOrigins = []string{config.Cfg.Common.Host}
	} else {
		// 测试环境下允许 localhost 的请求
		c.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	c.AllowCredentials = true
	return cors.New(c)
}
