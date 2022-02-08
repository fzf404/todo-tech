/*
 * @Author: fzf404
 * @Date: 2022-01-22 14:53:01
 * @LastEditTime: 2022-02-08 17:11:12
 * @Description: token 处理
 */
package utils

import (
	"time"
	"todo-list/config"
	"todo-list/model"

	"github.com/golang-jwt/jwt"
)

/**
 * @description: 签名密钥
 */
var jwtKey = []byte(config.Cfg.JWT.Key)

/**
 * @description: 分发 Token
 * @param {model.User} user
 */
func ReleaseToken(u model.User) (string, error) {

	// token 结构生成
	claims := &model.Claims{
		// 使用 ID、Username 作为有效载荷
		UID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + config.Cfg.JWT.Expires, // 超时时间
			NotBefore: time.Now().Unix() - 1000,                   // 签名生效时间
			Issuer:    config.Cfg.JWT.Issuer,                      // 发行人
		},
	}

	// 将 Claims 加密存储为 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 加入 key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/**
 * @description: 解析 Token
 * @param {string} tokenString
 */
func ParseToken(tokenString string) (*jwt.Token, *model.Claims, error) {
	claims := &model.Claims{}
	// 解码
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})
	return token, claims, err
}
