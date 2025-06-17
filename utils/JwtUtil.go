package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtExpireDuration = 7 * 24 * time.Hour // 7天

func GenerateToken(userId int64, secret string) (string, error) {
	claims := jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(JwtExpireDuration).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
func ParseToken(tokenString string, secret string) (int64, error) {
	// 解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是 HMAC（HS256）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	// 断言并提取 claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 提取 userId
		if userIdFloat, ok := claims["userId"].(float64); ok {
			return int64(userIdFloat), nil
		}
		return 0, jwt.ErrTokenInvalidClaims
	}

	return 0, jwt.ErrTokenInvalidClaims
}
