package util

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var JwtSecret string = "7879&2747"
var JwtExpireTime time.Duration = 24

var jwtSecret = []byte(JwtSecret) //jwt密钥

type Claims struct {
	ID uint
	jwt.StandardClaims
}

//生成Token
func GeterateToken(id uint, mobile string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(JwtExpireTime * time.Hour)
	claims := Claims{
		id,
		//mobile,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//校验和解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
