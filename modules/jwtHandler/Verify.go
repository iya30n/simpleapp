package jwtHandler

import "github.com/golang-jwt/jwt/v4"

func Verify(token string) bool {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return false
	}

	// TODO: check if username inside of the token does not exists on admins table, return err

	return tkn.Valid
}