package login

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

func ParseToken(token, tokenSecret string) (*jwt.RegisteredClaims, error) {
	tokenParsed, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return &jwt.RegisteredClaims{}, err
	}
	claims := tokenParsed.Claims.(*jwt.RegisteredClaims)
	return claims, nil
}
