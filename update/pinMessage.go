package update

type PinMessage struct {
	ID        string `json:"msgId,omitempty"`
	Chat      Chat   `json:"chat,omitempty"`
	User      User   `json:"from,omitempty"`
	Text      string `json:"text,omitempty"`
	Format    Format `json:"format,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type UnpinMessage struct {
	ID        string `json:"msgId,omitempty"`
	Chat      Chat   `json:"chat,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}
