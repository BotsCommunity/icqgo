package update

type Message struct {
	ID        string `json:"msgId,omitempty"`
	Chat      Chat   `json:"chat,omitempty"`
	User      User   `json:"from,omitempty"`
	Text      string `json:"text,omitempty"`
	Format    Format `json:"format,omitempty"`
	Parts     []Part `json:"parts,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}

type EditMessage struct {
	ID              string `json:"msgId,omitempty"`
	Chat            Chat   `json:"chat,omitempty"`
	User            User   `json:"from,omitempty"`
	Timestamp       int    `json:"timestamp,omitempty"`
	Text            string `json:"text,omitempty"`
	Format          Format `json:"format,omitempty"`
	EditedTimestamp int    `json:"editedTimestamp,omitempty"`
}

type DeleteMessage struct {
	ID        string `json:"msgId,omitempty"`
	Chat      Chat   `json:"chat,omitempty"`
	Timestamp int    `json:"timestamp,omitempty"`
}
