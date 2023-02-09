package jwtHandler

import (
	"simpleapp/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(admin models.Admin) (string, error) {
	expirationTime := time.Now().Add(30 * time.Second)

	claims := &Claims{
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(JwtKey)
}
