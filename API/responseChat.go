package API

type ResponseChat struct {
	Info
	Type           string `json:"type,omitempty"`
	Domain         string `json:"nick,omitempty"`
	Name           string `json:"firstName,omitempty"`
	Surname        string `json:"lastName,omitempty"`
	About          string `json:"about,omitempty"`
	IsBot          bool   `json:"isBot,omitempty"`
	Language       string `json:"language,omitempty"`
	Title          string `json:"title,omitempty"`
	Rules          string `json:"rules,omitempty"`
	InviteLink     string `json:"inviteLink,omitempty"`
	Public         bool   `json:"public,omitempty"`
	JoinModeration bool   `json:"joinModeration,omitempty"`
}

type ResponseChatAdmins struct {
	Info
	Admins []ChatMember `json:"admins,omitempty"`
}

type ResponseChatMembers struct {
	Info
	Members []ChatMember `json:"members,omitempty"`
}

type ResponseChatUsers struct {
	Users []ChatMember `json:"users,omitempty"`
}

type ChatMember struct {
	UserID  string `json:"userId,omitempty"`
	Creator bool   `json:"creator,omitempty"`
	Admin   bool   `json:"admin,omitempty"`
}
