package update

type CallbackQuery struct {
	ID      string          `json:"queryId,omitempty"`
	Chat    Chat            `json:"chat,omitempty"`
	User    User            `json:"from,omitempty"`
	Message CallbackMessage `json:"message,omitempty"`
	Data    string          `json:"callbackData,omitempty"`
}

type CallbackMessage struct {
	ID      int     `json:"eventId,omitempty"`
	Type    string  `json:"type,omitempty"`
	Payload Message `json:"payload,omitempty"`
}
