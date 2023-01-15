package utils

import (
	"github.com/dgrijalva/jwt-go"
	"tiktok/app/constant"
	"tiktok/app/schema"
)

var secret = []byte(constant.TOKEN_SECRET) // 签名密钥

// GenerateToken 生成 token
func GenerateToken(user schema.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = user.Id
	return token.SignedString(secret)
}

// VerifyToken 校验 token
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token.Claims.(jwt.MapClaims), nil
}
