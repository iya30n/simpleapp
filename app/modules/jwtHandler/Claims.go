package jwtHandler

import "github.com/golang-jwt/jwt/v4"

var JwtKey []byte = []byte("simple-app-jwt-key")

type Claims struct {
	Username string
	jwt.RegisteredClaims
}