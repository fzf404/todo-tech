/*
 * @Author: fzf404
 * @Date: 2022-01-22 22:27:30
 * @LastEditTime: 2022-02-05 20:26:35
 * @Description: 工具
 */

package utils

import (
	"reflect"
)

/**
* @description: 结构体反射赋值
 */
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() // 获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   // 获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
