/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:46:14
 * @LastEditTime: 2022-02-08 17:10:48
 * @Description: Todo
 */
package model

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

/**
 * @description: 标签
 */
type Tag []string

/**
 * @description: todos 数据表
 */
type Todo struct {
	gorm.Model
	UID      uint   // 用户 ID
	Title    string `json:"title"  binding:"min=1,max=64"`    // 文章标题
	Content  string `json:"content"  binding:"min=1,max=512"` // 内容
	Tag      Tag    `json:"tag"  binding:"required"`          // 标签
	Deadline uint   `json:"deadline"`                         // 截止时间
}

/**
 * @description: Todo Dto 数据
 */
type TodoDto struct {
	ID       uint     `json:"tid" binding:"required"`           // Todo ID
	Title    string   `json:"title" binding:"min=1,max=64"`     // 文章标题
	Content  string   `json:"content"  binding:"min=1,max=512"` // 内容
	Tag      []string `json:"tag" binding:"required"`           // 标签
	Deadline uint     `json:"deadline"`                         // 截止时间
}

/**
 * @description: 存入数据库前转为 string
 * @param {interface{}} value
 */
func (t *Tag) Scan(value interface{}) error {
	bytesValue, _ := value.([]byte)
	return json.Unmarshal(bytesValue, t)
}

/**
 * @description: 读出数据前转为 json
 */
func (t Tag) Value() (driver.Value, error) {
	return json.Marshal(t)
}

/**
 * @description: 删除 Todo 数据
 */
type TodoRemove struct {
	ID uint `json:"tid" binding:"required"`
}
