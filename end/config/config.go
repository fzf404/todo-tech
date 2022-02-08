/*
 * @Author: fzf404
 * @Date: 2022-01-22 13:52:00
 * @LastEditTime: 2022-02-08 17:09:14
 * @Description: 配置管理
 */

package config

import (
	"log"

	"github.com/spf13/viper"
)

/**
 * @description: 全局配置信息
 */
var Cfg Config

/**
 * @description: 配置结构
 */
type Config struct {
	Common Common `mapstructure:"common"`
	JWT    JWT    `mapstructure:"jwt"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Todo   Todo   `mapstructure:"todo"`
}

/**
 * @description: 项目配置
 */
type Common struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"` // 端口
}

/**
 * @description: 鉴权配置
 */
type JWT struct {
	Key     string `mapstructure:"key"`     // 签名密钥
	Expires int64  `mapstructure:"expires"` // 过期时间
	Issuer  string `mapstructure:"issuer"`  // 发行人
}

/**
 * @description: 数据库配置
 */
type Mysql struct {
	Username string `mapstructure:"username"` // 数据库用户名
	Password string `mapstructure:"password"` // 数据库密码
	Host     string `mapstructure:"host"`     // 服务器地址
	Port     string `mapstructure:"port"`     // 端口
	Database string `mapstructure:"database"` // 数据库名
	Config   string `mapstructure:"config"`   // 高级配置
}

type Todo struct {
	Page int `mapstructure:"page"` // 每页数据量
}

/**
 * @description: 数据库配置转换
 */
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}

/**
 * @description: 初始化配置文件
 */
func InitConfig() {
	// 文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// 文件位置 (文件夹)
	viper.AddConfigPath("./config")
	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}
	// 解析至结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatal("Unmarshal config failed: ", err)
	}
}
