/*
 * @Author: fzf404
 * @Date: 2022-01-22 13:56:31
 * @LastEditTime: 2022-02-05 20:07:35
 * @Description: 数据库管理
 */

package database

import (
	"log"
	"todo-list/config"
	"todo-list/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @description: 数据库全局变量
 */
var MysqlDB *gorm.DB

/**
 * @description: 初始化数据库
 */
func InitMysql() {
	// 获取数据库配置
	m := config.Cfg.Mysql
	// 连接
	if db, err := gorm.Open(mysql.Open(m.Dsn())); err != nil {
		log.Fatal("Connect mysql failed: ", err)
	} else {
		// 自动建表
		db.AutoMigrate(&model.User{})
		db.AutoMigrate(&model.Todo{})
		MysqlDB = db
	}
}
