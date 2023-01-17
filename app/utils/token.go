package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"tiktok/app/constant"
)

var secret = []byte(constant.TokenSecret) // 签名密钥

// GenerateToken 生成 token
func GenerateToken(userId int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims[constant.USERID] = userId
	return token.SignedString(secret)
}

// VerifyToken 校验 token
func VerifyToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return 1, err
	}

	if !token.Valid {
		return 1, jwt.ErrSignatureInvalid
	}

	userInfo := token.Claims.(jwt.MapClaims)[constant.USERID]
	currentUserId, err := strconv.ParseInt(fmt.Sprintf("%v", userInfo), 10, 64)
	return currentUserId, nil
}
