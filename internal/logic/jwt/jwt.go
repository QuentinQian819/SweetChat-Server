package jwt

import (
	"context"
	"errors"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明
type Claims struct {
	UserId uint64 `json:"user_id"`
	Phone  string `json:"phone"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT token
func GenerateToken(userId uint64, phone string) (string, error) {
	// 从配置中获取密钥和过期时间
	config := g.Cfg()
	secret := config.MustGet(context.Background(), "jwt.secret").String()
	expireHours := config.MustGet(context.Background(), "jwt.expire").Int64()

	nowTime := gtime.Now()
	expireTime := nowTime.Add(time.Duration(expireHours) * time.Second)

	claims := Claims{
		UserId: userId,
		Phone:  phone,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime.Time),
			IssuedAt:  jwt.NewNumericDate(nowTime.Time),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ParseToken 解析JWT token
func ParseToken(tokenString string) (*Claims, error) {
	config := g.Cfg()
	secret := config.MustGet(context.Background(), "jwt.secret").String()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
