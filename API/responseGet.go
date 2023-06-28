package API

type ResponseGet struct {
	Info
	UserID string `json:"userId,omitempty"`
	Domain string `json:"nick,omitempty"`
	Name   string `json:"firstName,omitempty"`
	About  string `json:"about,omitempty"`
	Photo  string `json:"photo,omitempty"`
}
