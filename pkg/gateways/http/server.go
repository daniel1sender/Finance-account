package http

const (
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

type Error struct {
	Reason string `json:"reason"`
}
