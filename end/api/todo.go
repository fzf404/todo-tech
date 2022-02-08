/*
 * @Author: fzf404
 * @Date: 2022-01-22 18:22:47
 * @LastEditTime: 2022-02-05 20:04:21
 * @Description: Todo 管理
 */
package api

import (
	"strconv"
	"todo-list/model"
	"todo-list/res"
	"todo-list/service"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 获取全部 Todo
 * @param {*gin.Context} c
 */
func AllTodo(c *gin.Context) {
	p := c.Query("page")         // 读取 query 参数
	page, err := strconv.Atoi(p) // 转为 int
	if err != nil {
		page = 1 // 如果未传入 & 出错，默认为 1
	}
	// 从中间件中获得用户信息
	uid, _ := c.Get("UserID")
	service.AllTodo(uid.(uint), page, c)
}

/**
 * @description: 搜索 Todo
 * @param {*gin.Context} c
 */
func SearchTodo(c *gin.Context) {
	title, ok := c.GetQuery("title") // 读取 query 参数
	if !ok || len(title) == 0 {
		res.FailWithMessage(res.ParamError, "提交信息非法", c)
		return
	}
	p := c.Query("page")         // 读取 query 分页参数
	page, err := strconv.Atoi(p) // 转为 int
	if err != nil {
		page = 1
	}
	uid, _ := c.Get("UserID") // 绑定用户信息
	service.SearchTodo(uid.(uint), title, page, c)
}

/**
 * @description: 增加 Todo
 * @param {*gin.Context} c
 */
func AddTodo(c *gin.Context) {
	// 绑定内容
	var t model.Todo
	if err := c.ShouldBind(&t); err != nil {
		res.FailWithDetailed(res.ParamError, err.Error(), "提交信息非法", c)
		return
	}

	uid, _ := c.Get("UserID") // 获取用户信息
	t.UID = uid.(uint)        // 写入 UID

	service.AddTodo(t, c)
}

/**
 * @description: 更新 Todo
 * @param {*gin.Context} c
 */
func UpdateTodo(c *gin.Context) {
	// 绑定信息
	var t model.TodoDto
	if err := c.ShouldBind(&t); err != nil {
		res.FailWithDetailed(res.ParamError, err.Error(), "提交信息非法", c)
		return
	}
	uid, _ := c.Get("UserID") // 获取用户信息
	service.UpdateTodo(uid.(uint), t, c)
}

/**
 * @description: 删除 Todo
 * @param {*gin.Context} c
 */
func RemoveTodo(c *gin.Context) {
	// 绑定信息
	var t model.TodoRemove
	if err := c.ShouldBind(&t); err != nil {
		res.FailWithDetailed(res.ParamError, err.Error(), "提交信息非法", c)
		return
	}

	uid, _ := c.Get("UserID") // 获取用户信息
	service.RemoveTodo(uid.(uint), t.ID, c)
}
