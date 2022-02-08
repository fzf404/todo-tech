/*
 * @Author: fzf404
 * @Date: 2022-01-22 12:56:50
 * @LastEditTime: 2022-02-05 19:49:24
 * @Description: 入口
 */
package main

import (
	"todo-list/config"
	"todo-list/database"
	"todo-list/router"
)

/**
 * @description: 初始化
 */
func init() {
	config.InitConfig()  // 初始化配置
	database.InitMysql() // 初始化数据库
}

/**
 * @description: 主函数
 */
func main() {
	r := router.InitRouter() // 初始化路由
	r.Run(":" + config.Cfg.Common.Port)
}
