package entities

import (
	"time"

	"github.com/google/uuid"
)

type Claims struct {
	TokenID     string
	Sub         string
	ExpTime     time.Time
	CreatedTime time.Time
}

func NewClaim(accountID string, expTime time.Duration) Claims {
	return Claims{
		TokenID:     uuid.NewString(),
		Sub:         accountID,
		ExpTime:     time.Now().Add(expTime),
		CreatedTime: time.Now().UTC(),
	}
}
