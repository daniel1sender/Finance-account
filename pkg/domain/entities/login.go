package entities

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	*jwt.StandardClaims
	origin_account_id string
	authorization     bool
}

func NewClaim(id string) Claims {
	claim := Claims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			IssuedAt:  time.Now().UTC().Unix(),
			Issuer:    "github.com/daniel1sender/Desafio-API",
		},
		origin_account_id: id,
		authorization:     true,
	}
	return claim
}
