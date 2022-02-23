package login

import (
	"fmt"

	"github.com/daniel1sender/Desafio-API/pkg/domain/entities"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id string) (string, error) {
	claim := entities.NewClaim(id)
	JWTtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	token, err := JWTtoken.SignedString(JWTtoken)
	if err != nil {
		return "", fmt.Errorf("error to get the signed token: %w", err)
	}
	return token, nil
}
