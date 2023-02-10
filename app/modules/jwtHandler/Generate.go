package jwtHandler

import (
	"simpleapp/app/models/Admin"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func Generate(admin Admin.AdminModel) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(JwtKey)
}
