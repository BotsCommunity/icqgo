package update

type Part struct {
	Type    string      `json:"type,omitempty"`
	Payload PartPayload `json:"payload,omitempty"`
}

type PartPayload struct {
	Mention
	File
	Forward
}

type Mention struct {
	UserID  string `json:"userId,omitempty"`
	Name    string `json:"firstName,omitempty"`
	Surname string `json:"lastName,omitempty"`
}

type File struct {
	FileId  string `json:"fileId,omitempty"`
	Type    string `json:"type,omitempty"`
	Caption string `json:"caption,omitempty"`
}

type Forward Message
