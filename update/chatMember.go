package update

type ChatMembers struct {
	Chat    Chat         `json:"chat,omitempty"`
	Members []ChatMember `json:"newMembers,omitempty"`
	AddedBy ChatMember   `json:"addedBy,omitempty"`
}

type LeftChatMembers struct {
	Chat      Chat         `json:"chat,omitempty"`
	Members   []ChatMember `json:"leftMembers,omitempty"`
	RemovedBy ChatMember   `json:"removedBy,omitempty"`
}

type ChatMember struct {
	ID      string `json:"userId,omitempty"`
	Name    string `json:"firstName,omitempty"`
	Surname string `json:"lastName,omitempty"`
}
