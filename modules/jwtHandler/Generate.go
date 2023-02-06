package jwtHandler

import (
	"simpleapp/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	username string `json:"username"`
	jwt.RegisteredClaims
}

var jwtKey []byte = []byte("simple-app-jwt-key")

func Generate(admin models.Admin) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(jwtKey)
}
