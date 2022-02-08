/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:47:41
 * @LastEditTime: 2022-02-08 17:10:38
 * @Description: Jwt
 */
package model

import (
	"github.com/golang-jwt/jwt"
)

// Claims jwt
type Claims struct {
	UID uint // 用户 ID
	jwt.StandardClaims
}
