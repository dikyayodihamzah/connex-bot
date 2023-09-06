package helper

import (
	"os"

	"github.com/dikyayodihamzah/connex-bot/model/domain"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	RoleId      string                  `json:"role_id"`
	Permissions []domain.RolePermission `json:"permissions"`
	jwt.StandardClaims
}

var secretKey = os.Getenv("SECRET_KEY")

func ParseJwt(cookie string) (Claims, error) {
	var claims Claims

	token, err := jwt.ParseWithClaims(cookie, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		return claims, err
	}

	return claims, err
}
