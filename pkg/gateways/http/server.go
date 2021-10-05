package http

const (
	JSONContentType = "application/json"
)

type Error struct {
	Reason string `json:"reason"`
}
