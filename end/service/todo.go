/*
 * @Author: fzf404
 * @Date: 2022-01-23 13:26:45
 * @LastEditTime: 2022-02-05 20:19:01
 * @Description: Todo 处理
 */
package service

import (
	"todo-list/config"
	"todo-list/database"
	"todo-list/model"
	"todo-list/res"
	"todo-list/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 获取全部 Todo
 * @param {uint} uid
 * @param {*gin.Context} c
 */
func AllTodo(uid uint, page int, c *gin.Context) {

	// 取得全部信息
	var raw []model.Todo
	database.MysqlDB.Where("uid = ?", uid).Limit(config.Cfg.Todo.Page).Offset(config.Cfg.Todo.Page * (page - 1)).Find(&raw)

	// 判断是否为空
	if len(raw) == 0 {
		res.FailWithMessage(res.NotFind, "什么都没有", c)
		return
	}

	// 信息格式化
	var data []model.TodoDto
	for _, t := range raw {
		var td model.TodoDto
		utils.StructAssign(&td, &t)
		td.ID = t.ID // 传入用户ID
		data = append(data, td)
	}

	res.OkWithData(data, c)
}

/**
 * @description: 搜索 Todo
 * @param {model.TodoSearch} s
 * @param {*gin.Context} c
 */
func SearchTodo(uid uint, t string, page int, c *gin.Context) {

	// 搜索信息
	var raw []model.Todo
	database.MysqlDB.Where("uid = ? and title LIKE ? ", uid, "%"+t+"%").Order("id desc").Find(&raw)

	// 为空判断
	if len(raw) == 0 {
		res.FailWithMessage(res.NotFind, "什么都没找到", c)
		return
	}

	// 信息格式化
	var data []model.TodoDto
	for _, t := range raw {
		var td model.TodoDto
		utils.StructAssign(&td, &t)
		td.ID = t.ID // 传入用户ID
		data = append(data, td)
	}

	res.OkWithData(data, c)
}

/**
 * @description: 增加 Todo
 * @param {model.Todo} n
 * @param {*gin.Context} c
 */
func AddTodo(t model.Todo, c *gin.Context) {
	// 插入 Todo
	database.MysqlDB.Create(&t)

	res.OkWithData(gin.H{"tid": t.ID}, c)
}

/**
 * @description: 更新 Todo
 * @param {model.Todo} n
 * @param {*gin.Context} c
 */
func UpdateTodo(uid uint, td model.TodoDto, c *gin.Context) {

	// 获取文章详情
	var todo model.Todo
	database.MysqlDB.First(&todo, td.ID)

	// 判断文章所有权
	if uid != todo.UID {
		res.FailWithMessage(res.NoAuth, "无权限", c)
		return
	}

	// 绑定信息
	utils.StructAssign(&todo, &td)

	// 保存
	database.MysqlDB.Save(&todo)

	res.OkWithData(gin.H{"tid": todo.ID}, c)
}

func RemoveTodo(uid uint, tid uint, c *gin.Context) {
	// 获取文章详情
	var todo model.Todo
	database.MysqlDB.First(&todo, tid)

	// 判断文章所有权
	if uid != todo.UID {
		res.FailWithMessage(res.NoAuth, "无权限", c)
		return
	}

	// 删除文章
	database.MysqlDB.Where("id = ?", tid).Delete(&model.Todo{})

	res.OkWithData(gin.H{"tid": todo.ID}, c)
}
