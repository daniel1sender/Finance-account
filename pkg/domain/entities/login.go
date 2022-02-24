package entities

import (
	"time"

	"github.com/google/uuid"
)

type Claims struct {
	TokenID   string
	Sub         string
	ExpTime     time.Time
	CreatedTime time.Time
}

func NewClaim(id string) Claims {
	return Claims{
		TokenID:   uuid.NewString(),
		Sub:         id,
		ExpTime:     time.Now().Add(time.Minute * 1),
		CreatedTime: time.Now().UTC(),
	}
}
