package update

import "encoding/json"

type Updates struct {
	OK      bool     `json:"ok,omitempty"`
	Updates []Update `json:"events,omitempty"`
}

type Update struct {
	ID      int             `json:"eventId,omitempty"`
	Type    string          `json:"type,omitempty"`
	Payload json.RawMessage `json:"payload,omitempty"`
}
