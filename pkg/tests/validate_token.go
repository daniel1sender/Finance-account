package tests

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func ValidateToken(tokenString string, accountID string, tokenSecret string) error {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return fmt.Errorf("expected no error but got '%v'", err)
	}
	claims := token.Claims.(*jwt.RegisteredClaims)
	if accountID != claims.Subject {
		return fmt.Errorf("expected account id equal sub")
	}
	if len(claims.ID) == 0 {
		return fmt.Errorf("expected not empty token id")
	}
	if !claims.VerifyExpiresAt(time.Now(), true) {
		return fmt.Errorf("expected non-zero 'expires at' time")
	}
	if !claims.VerifyIssuedAt(time.Now(), true) {
		return fmt.Errorf("expected non-zero 'issued at' time")
	}
	return nil
}
