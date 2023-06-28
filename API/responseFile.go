package API

type ResponseFile struct {
	Info
	MessageID string `json:"msgId,omitempty"`
	FileID    string `json:"fileId,omitempty"`
}
