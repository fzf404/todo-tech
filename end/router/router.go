/*
 * @Author: fzf404
 * @Date: 2022-01-22 15:22:05
 * @LastEditTime: 2022-02-05 20:20:17
 * @Description: 初始化路由
 */
package router

import (
	"todo-list/api"
	"todo-list/middleware"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 初始化路由
 */
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors()) // 跨域中间件

	route := r.Group("")
	{
		route.GET("/ping", api.Ping)
		route.POST("/user/login", api.Login)
		route.POST("/user/regist", api.Regist)

		route.Use(middleware.Auth())
		{
			route.GET("/user/verify", api.LoginVerify)
			route.GET("/todo/all", api.AllTodo)
			route.GET("/todo/search", api.SearchTodo)
			route.POST("/todo/add", api.AddTodo)
			route.POST("/todo/update", api.UpdateTodo)
			route.POST("/todo/remove", api.RemoveTodo)
		}
	}

	return r
}
