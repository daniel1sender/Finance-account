package entities

import (
	"time"

	"github.com/google/uuid"
)

type Claims struct {
	id          string
	sub         string
	expTime     int64
	createdTime int64
}

func NewClaim(id string) Claims {
	claim := Claims{
		id:          uuid.NewString(),
		sub:         id,
		expTime:     time.Now().Add(time.Minute * 1).Unix(),
		createdTime: time.Now().UTC().Unix(),
	}
	return claim
}
