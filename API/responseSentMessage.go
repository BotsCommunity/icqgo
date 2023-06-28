package API

type ResponseSentMessage struct {
	Info
	MessageID string `json:"msgId,omitempty"`
}
